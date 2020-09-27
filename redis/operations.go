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

func MigrateRedis(c *gin.Context) {
	// 获取请求参数
	containerName := c.Request.URL.Query().Get("container")
	checkpointID := c.Request.URL.Query().Get("checkpointID")
	destIP := c.Request.URL.Query().Get("destIP")
	destPort := c.Request.URL.Query().Get("destPort")
	checkpointDir := c.Request.URL.Query().Get("checkpointDir")
	// get all infos of a container
	containerJson, err := container.Inspect(containerName)
	if err != nil {
		fmt.Printf("container.Inspect err: %v\n", err)
		container.ReportErr(c, err)
		panic(err)
	}
	// get image name of the container to be migrated
	imageName, err := container.GetImageRepoTags(containerName)
	if err != nil {
		fmt.Printf("container.GetImageRepoTags err: %v\n", err)
		container.ReportErr(c, err)
		panic(err)
	}

	if checkpointDir == "" {
		//checkpointDir = migration.DefaultChkPDirPrefix + container.GetContainerFullID(containerName) + "/" + checkpointID
		checkpointDir = migration.DefaultChkPDirPrefix + containerJson.ID + "/" + checkpointID
	}

	// 1 send container create request
	cmd, err := json.Marshal(containerJson.Config.Cmd)
	cli := client.Cli{}
	createOpts := model.CreateOpts{
		ContainerName: containerName,
		ImageName:     imageName,
		HostPort:      "",
		ContainerPort: "",
		Cmd:           string(cmd),
		DestIP:        destIP,
	}

	_, err = cli.SendContainerCreate(createOpts)
	if err != nil {
		fmt.Printf("SendContainerCreate err: %v\n", err)
		container.ReportErr(c, err)
	}
	// 2 create a checkpoint
	chOpts := model.CheckpointOpts{
		CheckPointID:  checkpointID,
		CheckPointDir: checkpointDir,
	}
	err = container.CreateCheckpoint(c, chOpts)
	if err != nil {
		fmt.Printf("CreateCheckpoint err: %v\n", err)
		container.ReportErr(c, err)
		panic(err)
	}

	// 3 push checkpoint to destination node
	MigOpts := model.MigrationOpts{
		CheckpointOpts: chOpts,
		DestIP:         destIP,
		DestPort:       destPort,
	}
	err = migration.PushCheckpoint(MigOpts)
	if err != nil {
		fmt.Printf("Push Checkpoint err: %v\n", err)
		container.ReportErr(c, err)
		panic(err)
	}
	//

	c.JSON(200, gin.H{
		"result": "success",
	})
}
