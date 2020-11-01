package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/model"
	"github.com/yufeifly/migrator/utils"
)

func Run(c *gin.Context) {
	ImageName := c.Query("ImageName")
	HostPort := c.Query("HostPort")
	ContainerPort := c.Query("ContainerPort")

	runOpts := model.RunOpts{
		ImageName:     ImageName,
		HostPort:      HostPort,
		ContainerPort: ContainerPort,
	}

	err := container.RunContainer(runOpts)
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}
	c.JSON(200, gin.H{"docker run": "success"})
}
