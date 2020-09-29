package redis

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxyd/dal"
)

// Get redis get
func Get(c *gin.Context) {
	header := "redis.Get"
	key := c.Request.URL.Query().Get("key")

	val, err := dal.GetKV(key)
	if err != nil {
		logrus.Errorf("%s, err: %v", header, err)
		logrus.Panic(err)
	}
	c.JSON(200, gin.H{
		"value: ": val,
	})
}

// Set redis set
func Set(c *gin.Context) {
	key := c.Query("key")
	val := c.Query("val")

	logrus.WithFields(logrus.Fields{
		"key":   key,
		"value": val,
	}).Warn("about to set pair")

	err := dal.SetKV(key, val)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key":   key,
			"value": val,
		}).Error("set pair failed")
		logrus.Panic(err)
	}

	c.JSON(200, gin.H{
		"result": "success",
	})
}
