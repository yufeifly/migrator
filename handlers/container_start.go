package handlers

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/cusErr"
	"github.com/yufeifly/migrator/model"
	"github.com/yufeifly/migrator/utils"
)

// Start ContainerStart handler
func Start(c *gin.Context) {

	containerID := c.PostForm("ContainerID")
	checkpointID := c.PostForm("CheckpointID")
	checkpointDir := c.PostForm("CheckpointDir")

	if containerID == "" {
		utils.ReportErr(c, cusErr.ErrParamsNotValid)
		return
	}

	startOpts := model.StartOpts{
		ContainerID: containerID,
		CStartOpts: types.ContainerStartOptions{
			CheckpointID:  checkpointID,
			CheckpointDir: checkpointDir,
		},
	}

	err := container.StartContainer(startOpts)
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}
	c.JSON(200, gin.H{"result": "success"})
}
