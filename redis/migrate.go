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
		ContainerName: containerName,
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
	containerName := migrateOpts.ContainerName
	checkpointDir := migrateOpts.CheckpointDir
	destIP := migrateOpts.DestIP
	checkpointID := migrateOpts.CheckpointID
	destPort := migrateOpts.DestPort
	// get all infos of a container
	containerJson, err := container.Inspect(containerName)
	if err != nil {
		fmt.Printf("container.Inspect err: %v\n", err)
		return err
	}
	// get image name of the container to be migrated
	//imageName, err := container.GetImageRepoTags(containerName)
	imageName, err := container.GetImageByImageID(containerJson.Image)
	if err != nil {
		fmt.Printf("container.GetImageRepoTags err: %v\n", err)
		return err
	}

	if checkpointDir == "" {
		//checkpointDir = migration.DefaultChkPDirPrefix + container.GetContainerFullID(containerName) + "/" + checkpointID
		checkpointDir = migration.DefaultChkPDirPrefix + containerJson.ID
	}

	// 1 send container create request
	cmd, err := json.Marshal(containerJson.Config.Cmd)
	cmdStr := string(cmd)
	fmt.Printf("sender cmd: %v\n", cmdStr)

	cli := client.Cli{}
	createReqOpts := model.CreateReqOpts{
		CreateOpts: model.CreateOpts{
			ContainerName: containerName,
			ImageName:     imageName,
			HostPort:      "",
			ContainerPort: "",
			Cmd:           cmdStr,
		},
		DestIP:   destIP,
		DestPort: destPort,
	}

	rawResp, err := cli.SendContainerCreate(createReqOpts)
	if err != nil {
		fmt.Printf("SendContainerCreate err: %v\n", err)
		return err
	}
	var resp map[string]interface{}
	json.Unmarshal(rawResp, &resp)
	containerID := resp["containerId"].(string)
	fmt.Printf("create result: %v\n", containerID)

	// 2 create a checkpoint
	chOpts := model.CheckpointOpts{
		Container:     containerName,
		CheckPointID:  checkpointID,
		CheckPointDir: checkpointDir,
	}
	fmt.Printf("checkpoint opts : %v\n", chOpts)
	err = container.CreateCheckpoint(chOpts)
	if err != nil {
		fmt.Printf("CreateCheckpoint err: %v\n", err)
		return err
	}

	// 3 push checkpoint to destination node
	PushOpts := model.PushOpts{
		ContainerID:    containerID,
		CheckpointOpts: chOpts,
		DestIP:         destIP,
		DestPort:       destPort,
	}
	err = migration.PushCheckpoint(PushOpts)
	if err != nil {
		fmt.Printf("Push Checkpoint err: %v\n", err)
		return err
	}
	return nil
}
