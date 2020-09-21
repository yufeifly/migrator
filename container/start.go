package container

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
)



func Start(c *gin.Context) {
	containerId := c.Request.URL.Query().Get("containerId")
	err := cli.ContainerStart(context.Background(), containerId, types.ContainerStartOptions{})
	if err != nil {
		fmt.Println("start container failed")
		c.JSON(200, gin.H{
			"result": "failed",
		})
		panic(err)
	}
	c.JSON(200, gin.H{
		"result": "success",
	})
}
