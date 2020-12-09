package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/redis"
	"github.com/yufeifly/migrator/utils"
	"net/http"
)

// Get redis get handler
func Get(c *gin.Context) {
	header := "redis.Get"

	key := c.Query("key")
	serviceID := c.Query("service") // of worker

	val, err := redis.Get(serviceID, key)
	if err != nil {
		logrus.Errorf("%s, err: %v", header, err)
		utils.ReportErr(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, val)
	}
}

// Set redis set handler
func Set(c *gin.Context) {
	key := c.PostForm("key")
	val := c.PostForm("value")
	serviceID := c.PostForm("service")

	logrus.WithFields(logrus.Fields{
		"key":   key,
		"value": val,
	}).Debug("about to set pair")

	err := redis.Set(serviceID, key, val)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key":   key,
			"value": val,
		}).Error("set pair failed")
		utils.ReportErr(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}
