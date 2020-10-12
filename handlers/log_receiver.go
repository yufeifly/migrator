package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ReceiveLog(c *gin.Context) {
	var data []string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	logrus.Warnf("data to be consumed: %v", data)

	c.JSON(200, "log consumed")
}
