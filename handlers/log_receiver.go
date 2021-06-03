package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types/log"
	"github.com/yufeifly/migrator/task"
	"net/http"
)

func ReceiveLog(c *gin.Context) {
	var logWithCID log.LogWithCID

	if err := c.ShouldBindJSON(&logWithCID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	logrus.Warnf("logWithCID: %v", logWithCID)

	logJson, _ := json.Marshal(logWithCID.Log)

	q := task.DefaultMapper.GetTaskQueue(logWithCID.CID)
	if q == nil {
		q := task.NewQueue()
		task.DefaultMapper.AddTaskQueue(logWithCID.CID, q)
		logrus.Warn("ReceiveLog: new a task queue")
	}

	if task.DefaultMapper.GetTaskQueue(logWithCID.CID) != nil {
		logrus.Infof("push a log to queue, ProxyServiceID: %v", logWithCID.CID)
		task.DefaultMapper.GetTaskQueue(logWithCID.CID).Push(string(logJson)) // push a log to task queue
	} else {
		logrus.Panic("task.NewQueue failed")
	}

	c.JSON(http.StatusOK, "log received")
}
