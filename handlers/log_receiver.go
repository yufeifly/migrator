package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types/logger"
	"github.com/yufeifly/migrator/task"
	"net/http"
)

func ReceiveLog(c *gin.Context) {
	//ProxyServiceID := c.PostForm("Service")
	//logrus.Infof("ProxyServiceID: %v", ProxyServiceID)
	var logWithID logger.LogWithServiceID
	//var log model.Log
	if err := c.ShouldBindJSON(&logWithID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	logrus.Warnf("logWithID: %v", logWithID)

	logJson, _ := json.Marshal(logWithID.Log)

	q := task.DefaultMapper.GetTaskQueue(logWithID.ProxyServiceID)
	if q == nil {
		q := task.NewQueue()
		task.DefaultMapper.AddTaskQueue(logWithID.ProxyServiceID, q)
		logrus.Warn("ReceiveLog: new a task queue")
	}

	if task.DefaultMapper.GetTaskQueue(logWithID.ProxyServiceID) != nil {
		logrus.Infof("push a log to queue, ProxyServiceID: %v", logWithID.ProxyServiceID)
		task.DefaultMapper.GetTaskQueue(logWithID.ProxyServiceID).Push(string(logJson)) // push a log to task queue
	} else {
		logrus.Panic("task.NewQueue failed")
	}

	c.JSON(http.StatusOK, "log received")
}
