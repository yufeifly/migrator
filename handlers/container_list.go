package handlers

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/utils"
	"net/http"
)

// List handler for listing container(s)
func List(c *gin.Context) {
	header := "container.List"
	logrus.Infof("%s, list request: %v", header, c.Request)

	var listOpts types.ContainerListOptions
	if err := c.ShouldBindJSON(&listOpts); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	containers, err := container.ListContainers(listOpts)
	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panicf("%s, container.ListContainers panic: %v", header, err)
	}

	c.JSON(200, containers)
}
