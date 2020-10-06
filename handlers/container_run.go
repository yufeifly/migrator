package handlers

import "github.com/gin-gonic/gin"

func Run(c *gin.Context) {
	c.JSON(200, gin.H{
		"docker run": "to do",
	})
}
