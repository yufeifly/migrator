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
	err := cluster.LoadClusterConfig()
	if err != nil {
		logrus.Panicf("LoadClusterConfig failed, err: %v", err)
	}
	// init default scheduler
	scheduler.InitScheduler()
	//
	if !utils.IsDSTNode() {
		logrus.Info("PseudoRegistering")
		scheduler.PseudoRegister()
	}
	//
	r := gin.Default()
	router.InitRoutes(r)
	if err := r.Run(":6789"); err != nil {
		logrus.Errorf("gin.run err: %v", err)
	}
}
