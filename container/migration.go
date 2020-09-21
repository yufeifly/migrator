package container

import "github.com/gin-gonic/gin"

func Migrate(c *gin.Context) {
	c.JSON(200, gin.H{
		"migrate": "to do",
	})
}
