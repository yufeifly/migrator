package main

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/handlers"
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
	// redis operations
	// redis get func
	r.GET("/redis/get", handlers.Get)
	// redis set func
	r.POST("/redis/set", handlers.Set)
	// @deprecated redis migrate
	// r.POST("/redis/migration", redis.MigrateRedis)

	// container operations
	//  run a container
	r.POST("/docker/run", handlers.Run)
	//  start a container
	r.POST("/container/start", handlers.Start)
	//  list containers
	r.GET("/container/list", handlers.List)
	//  stop a container
	r.POST("/docker/stop", handlers.Stop)
	//  create a container
	r.POST("/docker/create", handlers.Create)
	//  create a container checkpoint
	r.POST("/docker/checkpoint/create", handlers.CheckpointCreate)
	// receive checkpoint and restore from it
	r.POST("/docker/checkpoint/restore", handlers.FetchCheckpointAndRestore)
	// push checkpoint to destination
	r.POST("/docker/checkpoint/push", handlers.CheckpointPush)
	// migrate a container
	r.POST("/docker/migrate", handlers.MigrateContainer)

	// listen and serve on 0.0.0.0:6789 (for windows "localhost:8080")
	if err := r.Run(":6789"); err != nil {
		logrus.Errorf("gin.run err: %v", err)
	}
}
