package migration

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types"
	"github.com/yufeifly/migrator/client"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/scheduler"
	"github.com/yufeifly/migrator/utils"
)

// MigrateOpts
type MigrateOpts struct {
	types.Address
	ServiceID     string
	ProxyService  string
	CheckpointID  string
	CheckpointDir string
}

// TryMigrate migrate redis service
func TryMigrate(mOpts MigrateOpts) error {
	header := "migration.TryMigrate"
	// get params
	ServiceID := mOpts.ServiceID         // real service id
	ProxyServiceID := mOpts.ProxyService // proxy id
	CheckpointID := mOpts.CheckpointID
	CheckpointDir := mOpts.CheckpointDir
	DestIP := mOpts.IP     // the destination ip
	DestPort := mOpts.Port // the destination port
	// get real service
	service, err := scheduler.Default().GetService(ServiceID)
	if err != nil {
		logrus.Errorf("%s, scheduler.DefaultScheduler.GetService err: %v", header, err)
		return err
	}
	// get all infos of a container
	containerJSON, err := container.Inspect(service.ContainerID)
	if err != nil {
		logrus.Errorf("%s, container.Inspect err: %v", header, err)
		return err
	}

	// get image name of the container to be migrated
	imageName := containerJSON.Config.Image

	// 1 send container create request
	// 1.1 get container's cmd in source node
	var CmdStr string
	if containerJSON.Config.Cmd != nil {
		cmd, err := json.Marshal(containerJSON.Config.Cmd)
		if err != nil {
			logrus.Errorf("%s, marshal cmd err: %v", header, err)
			return err
		}
		CmdStr = string(cmd)
		logrus.WithFields(logrus.Fields{
			"cmd": CmdStr,
		}).Debug("command of docker")
	}

	// 1.2 get container's port map in source node
	var PortBindingsStr string
	portBindings := containerJSON.HostConfig.PortBindings
	if portBindings != nil {
		pbJSON, err := json.Marshal(portBindings)
		if err != nil {
			logrus.Errorf("%s, marshal portBindings err: %v", header, err)
			return err
		}
		PortBindingsStr = string(pbJSON)
		logrus.WithFields(logrus.Fields{
			"PortBindings": PortBindingsStr,
		}).Debug("PortBindings")
	}

	var ExposedPortsStr string
	exposedPorts := containerJSON.Config.ExposedPorts
	if exposedPorts != nil {
		epJSON, err := json.Marshal(exposedPorts)
		if err != nil {
			return err
		}
		ExposedPortsStr = string(epJSON)
		logrus.WithFields(logrus.Fields{
			"ExposedPorts": ExposedPortsStr,
		}).Debug("ExposedPorts")
	}

	createReqOpts := types.CreateReqOpts{
		CreateOpts: types.CreateOpts{
			ContainerName: "", // empty string means a random name
			ImageName:     imageName,
			HostPort:      "", // empty string
			ContainerPort: "", // empty string
			PortBindings:  PortBindingsStr,
			ExposedPorts:  ExposedPortsStr,
			Cmd:           CmdStr,
		},
		Address: types.Address{
			IP:   DestIP,
			Port: DestPort,
		},
	}
	cli := client.NewClient(types.Address{
		IP:   DestIP,
		Port: DestPort,
	})
	rawResp, err := cli.SendContainerCreate(createReqOpts)
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
	containerID := resp["ContainerId"].(string) // the containerID of the created container in destination node
	logrus.WithFields(logrus.Fields{
		"ContainerID": containerID,
	}).Debug("container on dest node created")

	// 2 create a checkpoint
	// make the default checkpoint dir
	if CheckpointDir == "" {
		CheckpointDir = DefaultChkPDirPrefix + containerJSON.ID
	}
	chOpts := container.CheckpointReqOpts{
		Container:     service.ContainerID,
		CheckPointID:  CheckpointID,
		CheckPointDir: CheckpointDir,
	}

	err = container.CreateCheckpoint(chOpts)
	if err != nil {
		logrus.Errorf("%s, CreateCheckpoint err: %v", header, err)
		return err
	}

	// 3 push checkpoint to destination node
	PushOpts := PushOpts{
		CheckPointID:  chOpts.CheckPointID,
		CheckPointDir: chOpts.CheckPointDir,
		DestIP:        DestIP,
		DestPort:      DestPort,
		ContainerID:   containerID,                          // created in dst
		ServiceID:     utils.RenameService(mOpts.ServiceID), // make a name for dst service based on src service name
		ServicePort:   service.ServicePort,
		ProxyService:  ProxyServiceID,
	}
	err = PushCheckpoint(PushOpts)
	if err != nil {
		logrus.Errorf("%s, Push Checkpoint err: %v", header, err)
		return err
	}
	logrus.Warn("PushCheckpoint finished")
	return nil
}
