package migration

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxyd/container"
	"github.com/yufeifly/proxyd/model"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func CheckpointPush(c *gin.Context) {
	containerName := c.Request.URL.Query().Get("container")
	checkpointID := c.Request.URL.Query().Get("checkpointID")
	destIP := c.Request.URL.Query().Get("destIP")
	destPort := c.Request.URL.Query().Get("destPort")
	checkpointDir := c.Request.URL.Query().Get("checkpointDir")
	containerJson, err := container.Inspect(containerName)
	if err != nil {

	}
	if checkpointDir == "" {
		//checkpointDir = DefaultChkPDirPrefix + container.GetContainerFullID(containerName) + "/" + checkpointID
		checkpointDir = DefaultChkPDirPrefix + containerJson.ID + "/" + checkpointID
	}

	MigOpts := model.MigrationOpts{
		CheckpointOpts: model.CheckpointOpts{
			CheckPointID:  checkpointID,
			CheckPointDir: checkpointDir,
		},
		DestIP:   destIP,
		DestPort: destPort,
	}
	err = PushCheckpoint(MigOpts)
	if err != nil {
		c.JSON(200, gin.H{
			"result": "failed",
		})
		panic(err)
	}

	c.JSON(200, gin.H{
		"result": "success",
	})
}

// PushCheckpoint push checkpoint to destination and deliver restore request
func PushCheckpoint(migOpts model.MigrationOpts) error {
	// 这里应该与目的端有一个交互，以便知道目的端是否真实接收到检查点
	// 这里应该作为文件传输的客户端
	ip := migOpts.DestIP
	port := migOpts.DestPort

	urlPost := "http://" + ip + ":" + port + "/docker/checkpoint/restore"
	params := map[string]string{
		"CheckPointID":  migOpts.CheckPointID,
		"CheckPointDir": migOpts.CheckPointDir,
	}
	cpPath := migOpts.CheckPointDir

	files, err := getFilesFromCheckpoint(cpPath)
	if err != nil {
		fmt.Printf("PushCheckpoint err: %v\n", err)
		return err
	}

	//
	//for _, val := range files {
	//	fmt.Printf("filename: %v\n", val)
	//}
	//

	req, err := newFileUploadRequest(urlPost, cpPath, files, params)
	if err != nil {
		fmt.Printf("error to new upload file request:%s\n", err.Error())
		return err
	}
	//fmt.Println(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error to request to the server:%s\n", err.Error())
		return err
	}

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		fmt.Printf("error to request to the server:%s\n", err.Error())
		return err
	}

	defer resp.Body.Close()
	//fmt.Println(body)

	return nil
}

// getFilesFromCheckpoint get files from dir(pathname)
func getFilesFromCheckpoint(pathname string) ([]string, error) {
	var files []string
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return files, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			// ignore files in dir criu.work/restore-...
			if restoreDir(fi.Name()) {
				continue
			}

			adder, err := getFilesFromCheckpoint(pathname + "/" + fi.Name())
			if err != nil {
				return nil, err
			}
			for ind, val := range adder {
				adder[ind] = fi.Name() + "/" + val
			}
			files = append(files, adder...)
		} else {
			files = append(files, fi.Name())
		}
	}
	return files, nil
}

// newFileUploadRequest create a file upload request
func newFileUploadRequest(url string, cpPath string, paths []string, params map[string]string) (*http.Request, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for _, file := range paths {
		fullPath := cpPath + "/" + file
		// debug
		fmt.Printf("fullpath: %v\n", fullPath)
		fileP, err := os.Open(fullPath)
		if err != nil {
			return nil, err
		}

		part, err := writer.CreateFormFile(UploadFileKey, file)

		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, fileP)
		fileP.Close()
	}

	// 其他参数列表写入 body
	for k, v := range params {
		if err := writer.WriteField(k, v); err != nil {
			return nil, err
		}
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	return req, err
}

// restoreDir check if path is dir or not
func restoreDir(filename string) bool {
	return strings.HasPrefix(filename, "restore-")
}
