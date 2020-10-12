package handlers

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/utils"
)

// List handler for listing container(s)
func List(c *gin.Context) {
	header := "container.List"

	// if all=true or all=1 then docker ps -a
	var all bool
	allStr := c.Query("all")
	if allStr == "true" || allStr == "1" {
		all = true
	}

	containers, err := container.ListContainers(types.ContainerListOptions{
		All: all,
	})
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panicf("%s, container.ListContainers panic: %v", header, err)
	}

	c.JSON(200, containers)
}
