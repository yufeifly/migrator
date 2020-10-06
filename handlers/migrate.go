package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/migration"
	"github.com/yufeifly/migrator/model"
	"github.com/yufeifly/migrator/utils"
)

// MigrateRedis handler of migrating redis
func MigrateContainer(c *gin.Context) {
	// 获取请求参数
	containerName := c.Query("container")
	checkpointID := c.Query("checkpointID")
	destIP := c.Query("destIP")
	destPort := c.Query("destPort")
	checkpointDir := c.Query("checkpointDir")

	migrateOpts := model.MigrateOpts{
		Container:     containerName,
		CheckpointID:  checkpointID,
		CheckpointDir: checkpointDir,
		DestIP:        destIP,
		DestPort:      destPort,
	}
	err := migration.TryMigrate(migrateOpts)
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}
	//
	c.JSON(200, gin.H{
		"result": "success",
	})
}
