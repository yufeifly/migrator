package container

import "github.com/gin-gonic/gin"

func Stop(c *gin.Context) {
	c.JSON(200, gin.H{
		"docker stop": "to do",
	})
}
