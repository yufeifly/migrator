package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/migration"
	"github.com/yufeifly/migrator/utils"
	"net/http"
)

// MigrateRedis handler of migrating redis
func Migrate(c *gin.Context) {
	// 获取请求参数
	var migrateOpts migration.MigrateOpts
	if err := c.ShouldBindJSON(&migrateOpts); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	logrus.Debugf("MigrateContainer.MigrateOpts: %v", migrateOpts)
	err := migration.Migrate(migrateOpts)

	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panic(err)
	}
	logrus.Warn("migration.TryMigrate finished")
	//
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
