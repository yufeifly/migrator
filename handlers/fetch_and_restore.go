package handlers

import (
	"fmt"
	"github.com/yufeifly/migrator/api/types"
	"github.com/yufeifly/migrator/cluster"
	"net/http"
	"os"
	"time"

	ctypes "github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types/svc"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/scheduler"
	"github.com/yufeifly/migrator/task"
	"github.com/yufeifly/migrator/utils"
)

// ReceiveCheckpointAndRestore get checkpoint from source node and restore from it
func FetchCheckpointAndRestore(c *gin.Context) {
	header := "migration.FetchCheckpointAndRestore"
	// get params
	cpDir := c.PostForm("CheckPointDir")
	cpID := c.PostForm("CheckPointID")
	//cIDDst := c.PostForm("ContainerIDDest")
	cID := c.PostForm("ContainerID")
	sID := c.PostForm("ServiceID")
	servicePort := c.PostForm("ServicePort")

	// checkpoint path, example: /tmp/cp1
	cpPath := cpDir + "/" + cpID
	logrus.WithFields(logrus.Fields{
		"checkpoint path": cpPath,
		"checkpointID":    cpID,
	}).Info("the checkpoint path and ID received")

	// delete checkpoint dir if it exists
	if utils.FileExist(cpPath) {
		err := os.RemoveAll(cpPath)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "failed"})
			logrus.Panic(err)
		}
	}

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		logrus.Errorf("%s, get form err: %v", header, err)
		logrus.Panic(err)
	}
	files := form.File["files"]

	// create the dirs needed
	err = os.MkdirAll(cpPath+"/criu.work", 0766)
	if err != nil {
		logrus.Errorf("%s, make directory err: %v", header, err)
		logrus.Panic(err)
	}

	for _, file := range files {
		filename := cpPath + "/" + file.Filename
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("save uploaded file err: %s", err.Error()))
			logrus.Errorf("%s, save uploaded file err: %v", header, err)
			logrus.Panic(err)
		}
	}

	// start the container
	// 1 todo check if container created
	time.Sleep(1000 * time.Millisecond)
	// 2 start the container
	startOpts := container.StartOpts{
		ContainerID: cID,
		CStartOpts: ctypes.ContainerStartOptions{
			CheckpointID:  cpID,
			CheckpointDir: cpDir,
		},
	}
	logrus.Infof("startOpts: %v", startOpts)
	err = container.StartContainer(startOpts)
	if err != nil {
		logrus.Errorf("%s, start container err: %v", header, err)
	}

	// inform src node that it has started
	c.JSON(http.StatusOK, gin.H{"result": "success"})

	// register a redis service
	containerServ := scheduler.NewContainerServ(svc.ServiceOpts{
		CID:  cID,
		SID:  sID,
		Port: servicePort,
	})
	scheduler.Default().AddContainerServ(containerServ)
	logrus.Infof("%s, AddService finished, new service: %v", header, containerServ)

	// consume logs
	logrus.Warn("going to consume logs")
	srcNode := cluster.Node{
		Address: types.Address{
			IP:   c.ClientIP(),
			Port: "6789",
		},
	}
	go func(srcNode cluster.Node) {
		consumer := task.NewConsumer()
		err := consumer.Consume(cID, srcNode)
		if err != nil {
			logrus.Panic(err)
		}
		logrus.Info("consumer goroutine stopped")
	}(srcNode)
}
