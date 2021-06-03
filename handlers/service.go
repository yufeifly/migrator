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

// ServiceAdd add a container to service
func ServiceAdd(c *gin.Context) {
	ServiceID := c.PostForm("ServiceID")
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
		CID:  containerID,
		SID:  ServiceID,
		Port: servicePort,
	}

	scheduler.Default().AddContainerServ(scheduler.NewContainerServ(sOpts))

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
