package container

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxyd/utils"
	"time"
)

// stop a container
func Stop(c *gin.Context) {
	containerID := c.Request.URL.Query().Get("containerID")
	timeout := time.Second * 10

	err := cli.ContainerStop(ctx, containerID, &timeout)
	if err != nil {
		utils.ReportErr(c, err)
		panic(err)
	}

	logrus.WithFields(logrus.Fields{
		"ContainerID": containerID,
	}).Info("the container has been stopped")

	c.JSON(200, gin.H{
		"result": "success",
	})
}
