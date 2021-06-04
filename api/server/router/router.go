package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/migrator/handlers"
)

// InitRoutes init routes
func InitRoutes(r *gin.Engine) {
	// redis operations
	r.GET("/redis/get", handlers.Get)
	r.POST("/redis/set", handlers.Set)
	r.POST("/redis/delete", handlers.Delete)

	// container operations
	r.POST("/container/run", handlers.Run)
	r.POST("/container/start", handlers.Start)
	r.GET("/container/list", handlers.List)
	r.POST("/container/stop", handlers.Stop)
	r.POST("/container/create", handlers.Create)
	r.POST("/container/checkpoint/create", handlers.CheckpointCreate)
	// receive checkpoint and restore from it
	r.POST("/container/checkpoint/restore", handlers.FetchCheckpointAndRestore)
	r.POST("/migrate", handlers.Migrate)

	// logger
	r.POST("/logger", handlers.ReceiveLog)
	r.POST("/log/consume", handlers.LogConsumedAdder)

	// service
	r.POST("/service/add", handlers.ServiceAdd)
}
