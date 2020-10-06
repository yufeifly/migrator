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

	//containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
	//	All: all,
	//})
	containers, err := container.ListContainers(types.ContainerListOptions{
		All: all,
	})
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}

	list := make(gin.H)
	for _, Container := range containers {
		logrus.WithFields(logrus.Fields{
			"ContainerID": Container.ID[:10],
			"Image":       Container.Image,
		}).Infof("%s, List infos", header)
		list[Container.ID[:10]] = Container.Image
	}

	c.JSON(200, list)
}
