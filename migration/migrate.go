package migration

import (
	"encoding/json"
	"time"

	ctypes "github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types"
	"github.com/yufeifly/migrator/client"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/scheduler"
)

// Migrate migrate a whole service or a container
func Migrate(mOpts types.MigrateOpts) error {
	err := MigrateOneWithLogging(mOpts)
	if err != nil {
		return err
	}
	return nil
}

func MigrateOneWithLogging(options types.MigrateOpts) error {

	cServ, err := scheduler.Default().GetContainerServ(options.CID)
	if err != nil {
		logrus.Errorf("migration.MigrateOneWithLogging GetContainerServ failed, err: %v", err)
		return err
	}
	logrus.Debugf("migration.MigrateOneWithLogging service: %v", cServ)

	options.CID = cServ.CID

	// start logging
	cServ.Ticket().Logging()

	startedCh := make(chan struct{})
	// send migrate request to src node
	go func() {
		err := TryMigrate(options)
		if err != nil {
			logrus.Panicf("cli.SendMigrate failed, err: %v", err)
		}
		startedCh <- struct{}{}
	}()

	// write log files to dst
	// when dst starts, open redis connection
	// dst node consumes logs in the meantime
	// wait until all logs are consumed(no whole log file)
	ticker := time.NewTicker(100 * time.Microsecond)
FOR:
	for {
		select {
		case <-startedCh:
			logrus.Debug("migration.TryMigrateWithLogging, get value from chan(started)")
			sent := cServ.LogSent()
			if sent == 0 {
				logrus.Debug("migration.TryMigrateWithLogging, log sent number is 0, about to send the last log")
				break FOR
			}
		case <-ticker.C:
			if cServ.LoggingFinished() {
				logrus.Warn("migration.TryMigrateWithLogging, downtime start")
				cServ.Ticket().BanWrite()
				break FOR
			}
		case log := <-cServ.Logger().LogBuffer():
			err := cServ.SendLog(log, options.Address, false)
			if err != nil {
				logrus.Errorf("cServ.SendLog failed, err: %v", err)
				return err
			}
		}
	}

	// send the last log with flag "true" to dst,
	// true flag tells dst that this is the last one, so the consumer goroutine can stop
	err = cServ.SendLog(cServ.Logger().Log, options.Address, true)
	if err != nil {
		logrus.Errorf("migration.TryMigrateWithLogging, SendLastLog failed, err: %v", err)
		return err
	}

	// wait until the last log consumed by dst
	for {
		<-ticker.C
		if cServ.LoggingFinished() {
			logrus.Warn("migration.TryMigrateWithLogging, switching, requests redirect to dst node")
			// 1 inform the proxy the migration
			// 2 delete the container of the node
			break
		}
	}
	ticker.Stop()

	// downtime end, unset global lock
	cServ.Ticket().ReturnNormal()

	return nil
}

// TryMigrate migrate a container
func TryMigrate(mOpts types.MigrateOpts) error {
	header := "migration.TryMigrate"
	// get params
	CID := mOpts.CID
	SID := mOpts.SID
	CheckpointID := mOpts.CheckpointID
	CheckpointDir := mOpts.CheckpointDir
	// get container service
	containerServ, err := scheduler.Default().GetContainerServ(CID)
	if err != nil {
		logrus.Errorf("%s, scheduler.DefaultScheduler.GetcontainerServ err: %v", header, err)
		return err
	}
	// get all infos of a container
	cJSON, err := container.Inspect(containerServ.CID)
	if err != nil {
		logrus.Errorf("%s, container.Inspect err: %v", header, err)
		return err
	}

	createOptions, err := parseContainer(cJSON, mOpts.Address)
	if err != nil {
		logrus.Errorf("%s, parseContainer err: %v", header, err)
		return err
	}

	cli := client.NewClient(mOpts.Address)
	rawResp, err := cli.SendContainerCreate(createOptions)
	if err != nil {
		logrus.Errorf("%s, SendContainerCreate err: %v", header, err)
		return err
	}
	var resp map[string]interface{}
	err = json.Unmarshal(rawResp, &resp)
	if err != nil {
		logrus.Errorf("%s, Unmarshal response err: %v", header, err)
		return err
	}

	// make the default checkpoint dir
	if CheckpointDir == "" {
		CheckpointDir = DefaultChkPDirPrefix + cJSON.ID
	}
	chOpts := types.CheckpointReqOpts{
		Container:     containerServ.CID,
		CheckPointID:  CheckpointID,
		CheckPointDir: CheckpointDir,
	}

	err = container.CreateCheckpoint(chOpts)
	if err != nil {
		logrus.Errorf("%s, CreateCheckpoint err: %v", header, err)
		return err
	}

	// push checkpoint to destination node
	PushOpts := PushOpts{
		CheckPointID:  chOpts.CheckPointID,
		CheckPointDir: chOpts.CheckPointDir,
		Dest:          mOpts.Address,
		CID:           mOpts.CID,
		SID:           SID,
		Port:          containerServ.Port,
	}
	err = PushCheckpoint(PushOpts)
	if err != nil {
		logrus.Errorf("%s, Push Checkpoint err: %v", header, err)
		return err
	}
	logrus.Warn("PushCheckpoint finished")
	return nil
}

func parseContainer(cJSON ctypes.ContainerJSON, address types.Address) (types.CreateReqOpts, error) {
	// get image name of the container to be migrated
	imageName := cJSON.Config.Image

	// 1 get container's cmd in source node
	var cmdStr string
	if cJSON.Config.Cmd != nil {
		cmd, err := json.Marshal(cJSON.Config.Cmd)
		if err != nil {
			return types.CreateReqOpts{}, err
		}
		cmdStr = string(cmd)
		logrus.WithFields(logrus.Fields{
			"cmd": cmdStr,
		}).Debug("command of docker")
	}

	// 2 get container's port map in source node
	var portBindingsStr string
	portBindings := cJSON.HostConfig.PortBindings
	if portBindings != nil {
		pbJSON, err := json.Marshal(portBindings)
		if err != nil {
			return types.CreateReqOpts{}, err
		}
		portBindingsStr = string(pbJSON)
		logrus.WithFields(logrus.Fields{
			"PortBindings": portBindingsStr,
		}).Debug("PortBindings")
	}

	var exposedPortsStr string
	exposedPorts := cJSON.Config.ExposedPorts
	if exposedPorts != nil {
		epJSON, err := json.Marshal(exposedPorts)
		if err != nil {
			return types.CreateReqOpts{}, err
		}
		exposedPortsStr = string(epJSON)
		logrus.WithFields(logrus.Fields{
			"ExposedPorts": exposedPortsStr,
		}).Debug("ExposedPorts")
	}

	createOptions := types.CreateReqOpts{
		CreateOpts: types.CreateOpts{
			ContainerName: cJSON.Name, // empty string means a random name
			ImageName:     imageName,
			HostPort:      "", // empty string
			ContainerPort: "", // empty string
			PortBindings:  portBindingsStr,
			ExposedPorts:  exposedPortsStr,
			Cmd:           cmdStr,
		},
		Address: address,
	}
	return createOptions, nil
}
