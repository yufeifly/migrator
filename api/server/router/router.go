package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/migrator/handlers"
)

// InitRoutes init routes
func InitRoutes(r *gin.Engine) {
	// redis operations
	// redis get func
	r.GET("/redis/get", handlers.Get)
	// redis set func
	r.POST("/redis/set", handlers.Set)
	// redis delete func
	r.POST("/redis/delete", handlers.Delete)
	// @deprecated redis migrate
	// r.POST("/redis/migration", redis.MigrateRedis)

	// container operations
	//  run a container
	r.POST("/container/run", handlers.Run)
	//  start a container
	r.POST("/container/start", handlers.Start)
	//  list containers
	r.GET("/container/list", handlers.List)
	//  stop a container
	r.POST("/container/stop", handlers.Stop)
	//  create a container
	r.POST("/container/create", handlers.Create)
	//  create a container checkpoint
	r.POST("/container/checkpoint/create", handlers.CheckpointCreate)
	// receive checkpoint and restore from it
	r.POST("/container/checkpoint/restore", handlers.FetchCheckpointAndRestore)
	// push checkpoint to destination
	r.POST("/container/checkpoint/push", handlers.CheckpointPush)
	// migrate a container
	r.POST("/container/migrate", handlers.MigrateContainer)

	// logger
	r.POST("/logger", handlers.ReceiveLog)

	// service
	r.POST("/service/add", handlers.ServiceAdd)
}
