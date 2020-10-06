package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/migration"
	"github.com/yufeifly/migrator/model"
	"github.com/yufeifly/migrator/utils"
)

func CheckpointPush(c *gin.Context) {
	header := "migration.CheckpointPush"

	containerName := c.Request.URL.Query().Get("container")
	checkpointID := c.Request.URL.Query().Get("checkpointID")
	destIP := c.Request.URL.Query().Get("destIP")
	destPort := c.Request.URL.Query().Get("destPort")
	checkpointDir := c.Request.URL.Query().Get("checkpointDir")

	containerJson, err := container.Inspect(containerName)
	if err != nil {
		logrus.Errorf("%s, inspect container err: %v", header, err)
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}
	// get default dir to store checkpoint
	if checkpointDir == "" {
		checkpointDir = migration.DefaultChkPDirPrefix + containerJson.ID + "/" + checkpointID
	}

	PushOpts := model.PushOpts{
		CheckpointOpts: model.CheckpointOpts{
			CheckPointID:  checkpointID,
			CheckPointDir: checkpointDir,
		},
		DestIP:      destIP,
		DestPort:    destPort,
		ContainerID: containerJson.ID,
	}
	err = migration.PushCheckpoint(PushOpts)
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}

	c.JSON(200, gin.H{
		"result": "success",
	})
}
