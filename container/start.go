package container

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
)

func Start(c *gin.Context) {
	containerID := c.Request.URL.Query().Get("containerId")
	checkpointID := c.Request.URL.Query().Get("checkpointID")

	cStartOpts := types.ContainerStartOptions{
		CheckpointID: checkpointID,
	}

	err := cli.ContainerStart(context.Background(), containerID, cStartOpts)
	if err != nil {
		fmt.Println("start container failed")
		ReportErr(c, err)
		panic(err)
	}
	c.JSON(200, gin.H{
		"result": "success",
	})
}
