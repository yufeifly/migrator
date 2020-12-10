package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types/svc"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/scheduler"
	"github.com/yufeifly/migrator/utils"
	"net/http"
)

// ServiceAdd add a redis service
func ServiceAdd(c *gin.Context) {
	ServiceID := c.PostForm("ServiceID")
	ProxyServiceID := c.PostForm("ProxyServiceID")
	// choose a service port
	servicePort, err := utils.GetRandomPort()
	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panic(err)
	}
	// start a worker container
	rOpts := container.RunOpts{
		ImageName:     "docker.io/library/redis",
		HostPort:      servicePort,
		ContainerPort: servicePort,
	}
	containerID, err := container.RunContainer(rOpts)
	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panic(err)
	}

	sOpts := svc.ServiceOpts{
		ID:             ServiceID,
		ProxyServiceID: ProxyServiceID,
		ServicePort:    servicePort,
		Container:      containerID,
	}

	scheduler.Default().AddService(scheduler.NewService(sOpts))

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
