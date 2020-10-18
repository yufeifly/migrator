package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/model"
	"github.com/yufeifly/migrator/task"
	"net/http"
)

func ReceiveLog(c *gin.Context) {
	var log model.Log
	if err := c.ShouldBindJSON(&log); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	logrus.Warnf("data to be consumed: %v", log)
	serviceID := c.PostForm("Service")

	logJson, _ := json.Marshal(log)

	q := task.DefaultMapper.GetTaskQueue(serviceID)
	if q == nil {
		q := task.NewQueue()
		task.DefaultMapper.AddTaskQueue(serviceID, q)
	}

	if task.DefaultMapper.GetTaskQueue(serviceID) != nil {
		task.DefaultMapper.GetTaskQueue(serviceID).Push(string(logJson)) // push a log to task queue
	} else {
		logrus.Panic("task.NewQueue failed")
	}

	c.JSON(200, "log received")
}
