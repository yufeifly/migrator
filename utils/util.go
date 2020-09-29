package utils

import "github.com/gin-gonic/gin"

func ReportErr(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"result": err,
	})
}
