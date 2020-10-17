package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/redis"
)

// Get redis get
func Get(c *gin.Context) {
	header := "redis.Get"
	key := c.Query("key")
	serviceID := c.Query("service")

	val, err := redis.Get(serviceID, key)
	if err != nil {
		logrus.Errorf("%s, err: %v", header, err)
		c.JSON(200, gin.H{"failed: ": val})
	} else {
		c.JSON(200, val)
	}
}

// Set redis set
func Set(c *gin.Context) {
	key := c.PostForm("key")
	val := c.PostForm("value")
	serviceID := c.PostForm("service")

	logrus.WithFields(logrus.Fields{
		"key":   key,
		"value": val,
	}).Warn("about to set pair")

	err := redis.Set(serviceID, key, val)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key":   key,
			"value": val,
		}).Error("set pair failed")
		c.JSON(200, gin.H{"err": err})
	} else {
		c.JSON(200, gin.H{"result": "success"})
	}
}
