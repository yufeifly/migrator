package handlers

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/container"
	"github.com/yufeifly/migrator/model"
	"net/http"
	"os"
)

// ReceiveCheckpointAndRestore get checkpoint from source node and restore from it
func FetchCheckpointAndRestore(c *gin.Context) {
	header := "migration.FetchCheckpointAndRestore"

	cpDir := c.PostForm("CheckPointDir")
	cpID := c.PostForm("CheckPointID")
	cID := c.PostForm("ContainerID")
	cpPath := cpDir + "/" + cpID // example: /tmp/cp1

	logrus.WithFields(logrus.Fields{
		"checkpoint path": cpPath,
		"checkpointID":    cpID,
	}).Debug("the checkpoint path and ID received")
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

	c.JSON(200, gin.H{
		"result": "success",
	})
}
