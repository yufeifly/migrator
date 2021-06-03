package main

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/server/router"
	"github.com/yufeifly/migrator/cluster"
	"github.com/yufeifly/migrator/scheduler"
	"github.com/yufeifly/migrator/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	if utils.IsDebugEnabled() {
		utils.EnableDebug()
	}
}

func main() {
	// loading cluster
	err := cluster.LoadConfig()
	if err != nil {
		logrus.Panicf("load Cluster Config failed, err: %v", err)
	}
	// init default scheduler
	scheduler.Init()
	//
	if !utils.TargetNode() {
		logrus.Info("registering services")
		scheduler.RegisterServices()
	}
	//
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	router.InitRoutes(r)
	if err := r.Run(":6789"); err != nil {
		logrus.Errorf("gin.run err: %v", err)
	}
}
