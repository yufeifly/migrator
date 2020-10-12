package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/migration"
	"github.com/yufeifly/migrator/model"
	"github.com/yufeifly/migrator/utils"
	"net/http"
)

// MigrateRedis handler of migrating redis
func MigrateContainer(c *gin.Context) {
	// 获取请求参数
	var migrateOpts model.MigrateOpts
	if err := c.ShouldBindJSON(&migrateOpts); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	err := migration.TryMigrate(migrateOpts)
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}
	//
	c.JSON(200, gin.H{"result": "success"})
}
