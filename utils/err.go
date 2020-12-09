package utils

import "github.com/gin-gonic/gin"

func ReportErr(c *gin.Context, httpCode int, err error) {
	c.JSON(httpCode, gin.H{"result": err})
}
