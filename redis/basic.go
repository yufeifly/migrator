package redis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxyd/dal"
)

func Get(c *gin.Context) {
	key := c.Request.URL.Query().Get("key")

	val := dal.GetKV(key)
	c.JSON(200, gin.H{
		"value: ": val,
	})
}

func Set(c *gin.Context) {
	key := c.Request.URL.Query().Get("key")
	val := c.Request.URL.Query().Get("val")
	fmt.Printf("about to set key: %v, val: %v", key, val)

	dal.SetKV(key, val)

	c.JSON(200, gin.H{
		"result": "ok",
	})
}
