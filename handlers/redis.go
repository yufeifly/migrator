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
	key := c.Query("Key")
	cid := c.Query("ContainerID") // of worker

	val, err := redis.Get(cid, key)
	if err != nil {
		logrus.Errorf("%s, err: %v", "redis.Get", err)
		utils.ReportErr(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, val)
}

// Set redis set handler
func Set(c *gin.Context) {
	key := c.PostForm("Key")
	val := c.PostForm("Value")
	cid := c.PostForm("ContainerID")

	logrus.WithFields(logrus.Fields{
		"key":   key,
		"value": val,
	}).Debug("about to set pair")

	err := redis.Set(cid, key, val)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key":   key,
			"value": val,
		}).Error("set pair failed")
		utils.ReportErr(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}

func Delete(c *gin.Context) {
	key := c.PostForm("Key")
	cid := c.PostForm("ContainerID")

	err := redis.Delete(cid, key)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key": key,
		}).Error("delete pair failed")
		utils.ReportErr(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
