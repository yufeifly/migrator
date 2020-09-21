package container

import (
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	containerName := c.Request.URL.Query().Get("containerName")

	cCreateBody ,err := cli.ContainerCreate(ctx, &container.Config{}, &container.HostConfig{},
		&network.NetworkingConfig{}, containerName)
	if err != nil {
		fmt.Println("create container failed")
		c.JSON(200, gin.H{
			"result": "failed",
		})
		panic(err)
	}
	c.JSON(200, gin.H{
		"result": "success",
		"containerId": cCreateBody.ID,
	})
}
