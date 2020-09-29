package container

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// stop a container
func Stop(c *gin.Context) {
	containerID := c.Request.URL.Query().Get("containerId")
	timeout := time.Second * 10

	err := cli.ContainerStop(ctx, containerID, &timeout)
	if err != nil {
		ReportErr(c, err)
		panic(err)
	}

	fmt.Printf("container %v has been stopped", containerID)

	c.JSON(200, gin.H{
		"container stop": "success",
	})
}
