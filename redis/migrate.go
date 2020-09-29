package redis

import (
	"encoding/json"
	"fmt"
	"github.com/yufeifly/proxyd/client"
	"github.com/yufeifly/proxyd/model"

	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxyd/container"
	"github.com/yufeifly/proxyd/migration"
)

// MigrateRedis handler of migrating redis
func MigrateRedis(c *gin.Context) {
	// 获取请求参数
	containerName := c.Request.URL.Query().Get("container")
	checkpointID := c.Request.URL.Query().Get("checkpointID")
	destIP := c.Request.URL.Query().Get("destIP")
	destPort := c.Request.URL.Query().Get("destPort")
	checkpointDir := c.Request.URL.Query().Get("checkpointDir")

	migrateOpts := model.MigrateOpts{
		Container:     containerName,
		CheckpointID:  checkpointID,
		CheckpointDir: checkpointDir,
		DestIP:        destIP,
		DestPort:      destPort,
	}
	err := TryMigrate(migrateOpts)
	if err != nil {
		container.ReportErr(c, err)
		panic(err)
	}
	//
	c.JSON(200, gin.H{
		"result": "success",
	})
}

// TryMigrate migrate redis service
func TryMigrate(migrateOpts model.MigrateOpts) error {
	// get params
	Container := migrateOpts.Container // to identify container in source node
	CheckpointID := migrateOpts.CheckpointID
	CheckpointDir := migrateOpts.CheckpointDir
	DestIP := migrateOpts.DestIP     // the destination ip
	DestPort := migrateOpts.DestPort // the destination port

	// get all infos of a container
	containerJson, err := container.Inspect(Container)
	if err != nil {
		fmt.Printf("container.Inspect err: %v\n", err)
		return err
	}

	// get image name of the container to be migrated
	// imageName, err := container.GetImageByImageID(containerJson.Image)
	imageName := containerJson.Config.Image

	// make the default checkpoint dir
	if CheckpointDir == "" {
		//checkpointDir = migration.DefaultChkPDirPrefix + container.GetContainerFullID(containerName) + "/" + checkpointID
		CheckpointDir = migration.DefaultChkPDirPrefix + containerJson.ID
	}

	// 1 send container create request
	// 1.1 get container's cmd in source node
	cmd, err := json.Marshal(containerJson.Config.Cmd)
	if err != nil {

	}
	cmdStr := string(cmd)
	fmt.Printf("sender cmd: %v\n", cmdStr)

	// 1.2 get container's port map in source node
	var PortBindingsStr string
	portBindings := containerJson.HostConfig.PortBindings
	if portBindings != nil {
		pbJson, err := json.Marshal(portBindings)
		if err != nil {
			return err
		}
		PortBindingsStr = string(pbJson)
	}
	fmt.Printf("portBinding: %v\n", PortBindingsStr)

	var ExposedPortsStr string
	exposedPorts := containerJson.Config.ExposedPorts
	if exposedPorts != nil {
		epJson, err := json.Marshal(exposedPorts)
		if err != nil {
			return err
		}
		ExposedPortsStr = string(epJson)
	}
	fmt.Printf("exposedPorts: %v\n", ExposedPortsStr)

	cli := client.Cli{}
	createReqOpts := model.CreateReqOpts{
		CreateOpts: model.CreateOpts{
			ContainerName: "", // todo give dest container a nice name,empty string means a random name
			ImageName:     imageName,
			HostPort:      "",
			ContainerPort: "",
			PortBindings:  PortBindingsStr,
			ExposedPorts:  ExposedPortsStr,
			Cmd:           cmdStr,
		},
		DestIP:   DestIP,
		DestPort: DestPort,
	}

	rawResp, err := cli.SendContainerCreate(createReqOpts)
	if err != nil {
		fmt.Printf("SendContainerCreate err: %v\n", err)
		return err
	}
	var resp map[string]interface{}
	err = json.Unmarshal(rawResp, &resp)
	if err != nil {
		fmt.Printf("Unmarshal err: %v\n", err)
		return err
	}
	containerID := resp["containerId"].(string) // the containerID of the created container in destination node
	fmt.Printf("create result: %v\n", containerID)

	// 2 create a checkpoint
	chOpts := model.CheckpointOpts{
		Container:     Container,
		CheckPointID:  CheckpointID,
		CheckPointDir: CheckpointDir,
	}
	fmt.Printf("checkpoint opts : %v\n", chOpts)
	err = container.CreateCheckpoint(chOpts)
	if err != nil {
		fmt.Printf("CreateCheckpoint err: %v\n", err)
		return err
	}

	// 3 push checkpoint to destination node
	PushOpts := model.PushOpts{
		CheckpointOpts: chOpts,
		DestIP:         DestIP,
		DestPort:       DestPort,
		ContainerID:    containerID,
	}
	err = migration.PushCheckpoint(PushOpts)
	if err != nil {
		fmt.Printf("Push Checkpoint err: %v\n", err)
		return err
	}
	return nil
}
