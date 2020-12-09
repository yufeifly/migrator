package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/utils"
	"net/http"
	"strconv"
	"time"
)

// stop a container
func Stop(c *gin.Context) {
	ContainerID := c.PostForm("ContainerID")
	Seconds := c.PostForm("Timeout")

	var timeout *time.Duration
	if Seconds != "" {
		valSeconds, err := strconv.Atoi(Seconds)
		if err != nil {
			logrus.Panic(err)
		}
		timeoutCVal := time.Duration(valSeconds) * time.Second
		timeout = &timeoutCVal
	}

	//err := cli.ContainerStop(ctx, ContainerID, &timeout)
	err := container.StopContainer(ContainerID, timeout)
	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		panic(err)
	}

	logrus.WithFields(logrus.Fields{
		"ContainerID": ContainerID,
	}).Info("the container has been stopped")

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
