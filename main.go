package main

import (
	"fmt"

	"github.com/yufeifly/proxyd/container"
	"github.com/yufeifly/proxyd/dal"

	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/redis/get", func(c *gin.Context) {
		key := c.Request.URL.Query().Get("key")
		val := dal.GetKV(key)
		c.JSON(200, gin.H{
			"value: ": val,
		})
	})
	r.POST("/redis/set", func(c *gin.Context) {
		key := c.Request.URL.Query().Get("key")
		val := c.Request.URL.Query().Get("val")
		fmt.Printf("about to set key: %v, val: %v", key, val)
		dal.SetKV(key, val)
		c.JSON(200, gin.H{
			"result": "ok",
		})
	})

	// todo run a container
	r.POST("/docker/run", container.Run)

	// todo start a container
	r.POST("/docker/start", container.Start)

	// todo list containers
	r.GET("/docker/list", container.List)

	// todo stop a container
	r.POST("/docker/stop", container.Stop)

	// todo create a container
	r.POST("/docker/create", container.Create)

	// todo implement docker migration
	r.POST("/docker/migrate", container.Migrate)

	r.Run(":6789") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
