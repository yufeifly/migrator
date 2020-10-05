package main

import (
	"github.com/yufeifly/migrator/migration"
	"github.com/yufeifly/migrator/redis"
	"github.com/yufeifly/migrator/utils"

	"github.com/gin-gonic/gin"
	"github.com/yufeifly/migrator/container"
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
	r.GET("/redis/get", redis.Get)
	// redis set func
	r.POST("/redis/set", redis.Set)
	// @deprecated redis migrate
	r.POST("/redis/migration", redis.MigrateRedis)

	// container operations
	//  run a container
	r.POST("/docker/run", container.Run)
	//  start a container
	r.POST("/docker/start", container.Start)
	//  list containers
	r.GET("/docker/list", container.List)
	//  stop a container
	r.POST("/docker/stop", container.Stop)
	//  create a container
	r.POST("/docker/create", container.Create)
	//  create a container checkpoint
	r.POST("/docker/checkpoint/create", container.CheckpointCreate)
	// receive checkpoint and restore from it
	r.POST("/docker/checkpoint/restore", migration.FetchCheckpointAndRestore)
	// push checkpoint to destination
	r.POST("/docker/checkpoint/push", migration.CheckpointPush)

	// migrate a container
	r.POST("/migrate", migration.MigrateContainer)

	r.Run(":6789") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
