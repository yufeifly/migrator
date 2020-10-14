package migration

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/client"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/model"
)

// TryMigrate migrate redis service
func TryMigrate(migrateOpts model.MigrateOpts) error {
	header := "migration.TryMigrate"
	// get params
	Container := migrateOpts.Container // to identify container in source node
	CheckpointID := migrateOpts.CheckpointID
	CheckpointDir := migrateOpts.CheckpointDir
	DestIP := migrateOpts.IP     // the destination ip
	DestPort := migrateOpts.Port // the destination port

	// get all infos of a container
	containerJson, err := container.Inspect(Container)
	if err != nil {
		logrus.Errorf("%s, inspect err: %v", header, err)
		return err
	}

	// get image name of the container to be migrated
	imageName := containerJson.Config.Image

	// make the default checkpoint dir
	if CheckpointDir == "" {
		CheckpointDir = DefaultChkPDirPrefix + containerJson.ID
	}

	// 1 send container create request
	// 1.1 get container's cmd in source node
	var CmdStr string
	if containerJson.Config.Cmd != nil {
		cmd, err := json.Marshal(containerJson.Config.Cmd)
		if err != nil {
			logrus.Errorf("%s, marshal cmd err: %v", header, err)
			return err
		}
		CmdStr = string(cmd)
		logrus.WithFields(logrus.Fields{
			"cmd": CmdStr,
		}).Debug("command to send")
	}

	// 1.2 get container's port map in source node
	var PortBindingsStr string
	portBindings := containerJson.HostConfig.PortBindings
	if portBindings != nil {
		pbJson, err := json.Marshal(portBindings)
		if err != nil {
			return err
		}
		PortBindingsStr = string(pbJson)
		logrus.WithFields(logrus.Fields{
			"PortBindings": PortBindingsStr,
		}).Debug("PortBindings")
	}

	var ExposedPortsStr string
	exposedPorts := containerJson.Config.ExposedPorts
	if exposedPorts != nil {
		epJson, err := json.Marshal(exposedPorts)
		if err != nil {
			return err
		}
		ExposedPortsStr = string(epJson)
		logrus.WithFields(logrus.Fields{
			"ExposedPorts": ExposedPortsStr,
		}).Debug("ExposedPorts")
	}

	//cli := client.Client{}
	cli := client.NewClient()
	createReqOpts := model.CreateReqOpts{
		CreateOpts: model.CreateOpts{
			ContainerName: "", // todo give dest container a nice name,empty string means a random name
			ImageName:     imageName,
			HostPort:      "", // empty string
			ContainerPort: "", // empty string
			PortBindings:  PortBindingsStr,
			ExposedPorts:  ExposedPortsStr,
			Cmd:           CmdStr,
		},
		DestIP:   DestIP,
		DestPort: DestPort,
	}

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
	containerID := resp["containerId"].(string) // the containerID of the created container in destination node
	logrus.WithFields(logrus.Fields{
		"ContainerID": containerID,
	}).Debug("container on dest node created")

	// 2 create a checkpoint
	chOpts := model.CheckpointOpts{
		Container:     Container,
		CheckPointID:  CheckpointID,
		CheckPointDir: CheckpointDir,
	}

	err = container.CreateCheckpoint(chOpts)
	if err != nil {
		logrus.Errorf("%s, CreateCheckpoint err: %v", header, err)
		return err
	}

	// 3 push checkpoint to destination node
	PushOpts := model.PushOpts{
		CheckpointOpts: chOpts,
		DestIP:         DestIP,
		DestPort:       DestPort,
		ContainerID:    containerID,
	}
	err = PushCheckpoint(PushOpts)
	if err != nil {
		logrus.Errorf("%s, Push Checkpoint err: %v", header, err)
		return err
	}
	logrus.Warn("PushCheckpoint finished")
	return nil
}
