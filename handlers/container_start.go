package handlers

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/model"
	"github.com/yufeifly/migrator/utils"
)

// Start ContainerStart handler
func Start(c *gin.Context) {
	containerID := c.Query("containerId")
	checkpointID := c.Query("checkpointID")
	checkpointDir := c.Query("checkpointDir")

	startOpts := model.StartOpts{
		CStartOpts: types.ContainerStartOptions{
			CheckpointID:  checkpointID,
			CheckpointDir: checkpointDir,
		},
		ContainerID: containerID,
	}

	err := container.StartContainer(startOpts)
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}
	c.JSON(200, gin.H{
		"result": "success",
	})
}
