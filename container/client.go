package container

import (
	"context"
	"github.com/gin-gonic/gin"

	"github.com/docker/docker/client"
)

var (
	cli *client.Client
)

var ctx = context.Background()

func init() {
	var err error
	cli, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}
}

func ReportErr(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"result": err,
	})
}
