package container

import (
	"fmt"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	header := "[container:List]"

	var all bool
	allStr := c.Request.URL.Query().Get("all")
	if allStr == "true" {
		all = true
	} else {
		allInt, err := strconv.Atoi(allStr)
		if err != nil {
			fmt.Printf("%v, %v", header, err)
			c.JSON(200, gin.H{
				"result": "bad param",
			})
		}
		if allInt == 1 {
			all = true
		}
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: all,
	})
	if err != nil {
		c.JSON(200, gin.H{
			"result": "failed",
		})
		panic(err)
	}

	list := make(gin.H)
	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
		list[container.ID[:10]] = container.Image
	}

	c.JSON(200, list)
}
