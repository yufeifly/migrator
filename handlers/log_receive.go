package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/migrator/api/types/log"
	"github.com/yufeifly/migrator/task"
	"net/http"
)

func ReceiveLog(c *gin.Context) {
	var logWithCID log.LogWithCID

	if err := c.ShouldBindJSON(&logWithCID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	t := task.Default().GetTask(logWithCID.CID)
	if t == nil {
		newt := task.NewTask(logWithCID.CID)
		task.Default().AddTask(logWithCID.CID, newt)
	}

	task.Default().GetTask(logWithCID.CID).Push(logWithCID.Log)

	c.JSON(http.StatusOK, gin.H{
		"log received: ": logWithCID,
	})
}
