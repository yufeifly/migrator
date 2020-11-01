package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/model"
	"github.com/yufeifly/migrator/utils"
)

// CheckpointCreate handler for create a checkpoint for a container
func CheckpointCreate(c *gin.Context) {
	Container := c.Query("container")
	CheckpointID := c.Query("checkpointID")
	CheckpointDIR := c.Query("checkpointDIR")

	cpOpts := model.CheckpointOpts{
		Container:     Container,
		CheckPointID:  CheckpointID,
		CheckPointDir: CheckpointDIR,
	}
	if err := container.CreateCheckpoint(cpOpts); err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}

	c.JSON(200, gin.H{"migrate": "success"})
}
