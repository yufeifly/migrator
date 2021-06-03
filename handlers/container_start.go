package handlers

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/cuserr"
	"github.com/yufeifly/migrator/utils"
	"net/http"
)

// Start ContainerStart handler
func Start(c *gin.Context) {

	containerID := c.PostForm("ContainerID")
	checkpointID := c.PostForm("CheckpointID")
	checkpointDir := c.PostForm("CheckpointDir")

	if containerID == "" {
		utils.ReportErr(c, http.StatusBadRequest, cuserr.ErrParamsNotValid)
		return
	}

	startOpts := container.StartOpts{
		ContainerID: containerID,
		CStartOpts: types.ContainerStartOptions{
			CheckpointID:  checkpointID,
			CheckpointDir: checkpointDir,
		},
	}

	err := container.StartContainer(startOpts)
	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
