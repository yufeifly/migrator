package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/utils"
	"net/http"
)

// CheckpointCreate handler for create a checkpoint for a container
func CheckpointCreate(c *gin.Context) {
	Container := c.PostForm("Container")
	CheckpointID := c.PostForm("CheckpointID")
	CheckpointDIR := c.PostForm("CheckpointDIR")

	cpOpts := container.CheckpointReqOpts{
		Container:     Container,
		CheckPointID:  CheckpointID,
		CheckPointDir: CheckpointDIR,
	}
	if err := container.CreateCheckpoint(cpOpts); err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"CheckpointCreate": "success"})
}
