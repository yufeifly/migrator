package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/scheduler"
	"github.com/yufeifly/migrator/utils"
	"net/http"
)

// LogConsumedAdder ...
func LogConsumedAdder(c *gin.Context) {

	cID := c.PostForm("ContainerID")

	containerServ, err := scheduler.Default().GetContainerServ(cID)
	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panic(err)
	}
	containerServ.ConsumedAdder()

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
