package main

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/server/router"
	"github.com/yufeifly/migrator/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	if utils.IsDebugEnabled() {
		utils.EnableDebug()
	}
}

func main() {
	r := gin.Default()
	router.InitRoutes(r)
	if err := r.Run(":6789"); err != nil {
		logrus.Errorf("gin.run err: %v", err)
	}
}
