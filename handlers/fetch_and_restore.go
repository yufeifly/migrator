package handlers

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/model"
	"github.com/yufeifly/migrator/scheduler"
	"github.com/yufeifly/migrator/task"
	"github.com/yufeifly/migrator/utils"
	"net/http"
	"os"
	"time"
)

// ReceiveCheckpointAndRestore get checkpoint from source node and restore from it
func FetchCheckpointAndRestore(c *gin.Context) {
	header := "migration.FetchCheckpointAndRestore"
	// get params
	cpDir := c.PostForm("CheckPointDir")
	cpID := c.PostForm("CheckPointID")
	cID := c.PostForm("ContainerID")         // of dst
	serviceID := c.PostForm("ServiceID")     // of dst worker
	servicePort := c.PostForm("ServicePort") // of src worker, also dst worker
	proxyServiceID := c.PostForm("ProxyServiceID")

	// checkpoint path
	cpPath := cpDir + "/" + cpID // example: /tmp/cp1
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
	time.Sleep(500 * time.Millisecond)
	// 2 start the container
	startOpts := model.StartOpts{
		CStartOpts: types.ContainerStartOptions{
			CheckpointID:  cpID,
			CheckpointDir: cpDir,
		},
		ContainerID: cID,
	}
	err = container.StartContainer(startOpts)
	if err != nil {
		logrus.Errorf("%s, start container err: %v", header, err)
	}

	// inform proxy it has started. request: proxy -> src -> dst, so the respond: dst -> src -> proxy
	c.JSON(http.StatusOK, gin.H{"result": "success"})

	// register a redis service
	service := scheduler.NewService(model.ServiceOpts{
		ID:             serviceID,
		ProxyServiceID: proxyServiceID,
		ServicePort:    servicePort,
		Container:      cID,
	})
	scheduler.DefaultScheduler.AddService(service)
	logrus.Infof("%s, AddService finished, new service: %v", header, service)

	// consume logs
	// todo but which log belongs to it?
	logrus.Warn("going to consume logs")
	go func() {
		consumer := task.NewConsumer()
		err := consumer.Consume(proxyServiceID)
		if err != nil {
			logrus.Panic(err)
		}
		logrus.Info("consumer goroutine stopped")
	}()
}
