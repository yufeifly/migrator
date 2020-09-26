package migration

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxyd/container"
	"github.com/yufeifly/proxyd/model"
	"net/http"
	"os"
)

// ReceiveCheckpointAndRestore get checkpoint from source node and restore from it
func FetchCheckpointAndRestore(c *gin.Context) {
	// 文件传输的服务器端
	//cpDir := c.Request.URL.Query().Get("cpDir")
	cpDir := c.PostForm("CheckPointDir")
	cpID := c.PostForm("CheckPointID")
	// debug
	//fmt.Printf("cpDir: %v, cpID: %v\n", cpDir, cpID)
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["files"]

	// create the dirs needed
	err = os.MkdirAll(cpDir+"/criu.work", 0766)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		//fmt.Printf("file.filename: %v\n", file.Filename)
		//filename := cpDir + "/test/" + filepath.Base(file.Filename)
		filename := cpDir + "/" + file.Filename
		//fmt.Printf("filename: %v\n", filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}

	//c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files.", len(files)))
	// 启动容器
	// 1 check if container created
	// 2 start the container
	startOpts := model.StartOpts{
		CStartOpts: types.ContainerStartOptions{
			CheckpointID:  cpID,
			CheckpointDir: cpDir,
		},
		ContainerID: containerID,
	}
	container.StartContainer(startOpts)

	fmt.Printf("checkpointID: %v\n", cpID)

	c.JSON(200, gin.H{
		"result": "success",
	})
}
