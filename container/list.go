package container

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxyd/utils"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
)

// List handler for listing container(s)
func List(c *gin.Context) {
	header := "container.List"

	// if all=true or all=1 then docker ps -a
	var all bool
	allStr := c.Request.URL.Query().Get("all")
	if allStr == "true" {
		all = true
	} else {
		allInt, err := strconv.Atoi(allStr)
		if err != nil {
			logrus.Errorf("%s, err: %v", header, err)
			utils.ReportErr(c, err)
		}
		if allInt == 1 {
			all = true
		}
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: all,
	})
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}

	list := make(gin.H)
	for _, container := range containers {
		logrus.WithFields(logrus.Fields{
			"ContainerID": container.ID[:10],
			"Image":       container.Image,
		}).Info("List infos")
		list[container.ID[:10]] = container.Image
	}

	c.JSON(200, list)
}
