package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/utils"
	"net/http"
)

func Run(c *gin.Context) {
	ImageName := c.Query("ImageName")
	HostPort := c.Query("HostPort")
	ContainerPort := c.Query("ContainerPort")

	runOpts := container.RunOpts{
		ImageName:     ImageName,
		HostPort:      HostPort,
		ContainerPort: ContainerPort,
	}

	containerID, err := container.RunContainer(runOpts)
	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"containerID": containerID})
}
