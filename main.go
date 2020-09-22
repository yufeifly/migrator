package main

import (
	"github.com/yufeifly/proxyd/redis"

	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxyd/container"
)

func init() {

}

func main() {
	r := gin.Default()

	r.GET("/redis/get", redis.Get)
	r.POST("/redis/set", redis.Set)
	r.POST("/redis/migration", redis.MigrateRedis)

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

	r.Run(":6789") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
