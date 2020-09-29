package migration

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxyd/utils"

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
	header := "migration.CheckpointPush"

	containerName := c.Request.URL.Query().Get("container")
	checkpointID := c.Request.URL.Query().Get("checkpointID")
	destIP := c.Request.URL.Query().Get("destIP")
	destPort := c.Request.URL.Query().Get("destPort")
	checkpointDir := c.Request.URL.Query().Get("checkpointDir")

	containerJson, err := container.Inspect(containerName)
	if err != nil {
		logrus.Errorf("%s, inspect container err: %v", header, err)
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}
	// get default dir to store checkpoint
	if checkpointDir == "" {
		checkpointDir = DefaultChkPDirPrefix + containerJson.ID + "/" + checkpointID
	}

	PushOpts := model.PushOpts{
		CheckpointOpts: model.CheckpointOpts{
			CheckPointID:  checkpointID,
			CheckPointDir: checkpointDir,
		},
		DestIP:      destIP,
		DestPort:    destPort,
		ContainerID: containerJson.ID,
	}
	err = PushCheckpoint(PushOpts)
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}

	c.JSON(200, gin.H{
		"result": "success",
	})
}

// PushCheckpoint push checkpoint to destination and deliver restore request
func PushCheckpoint(migOpts model.PushOpts) error {
	header := "migration.PushCheckpoint"

	ip := migOpts.DestIP
	port := migOpts.DestPort

	urlPost := "http://" + ip + ":" + port + "/docker/checkpoint/restore"
	params := map[string]string{
		"ContainerID":   migOpts.ContainerID,
		"CheckPointID":  migOpts.CheckPointID,
		"CheckPointDir": migOpts.CheckPointDir + "/test",
	}
	cpPath := migOpts.CheckPointDir + "/" + migOpts.CheckPointID

	files, err := getFilesFromCheckpoint(cpPath)
	if err != nil {
		logrus.Errorf("%s, getFilesFromCheckpoint err: %v", header, err)
		return err
	}

	req, err := newFileUploadRequest(urlPost, cpPath, files, params)
	if err != nil {
		logrus.Errorf("%s, new file upload request err: %v", header, err)
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("%s, do request to server err: %v", header, err)
		return err
	}

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		logrus.Errorf("%s, read response body err: %v", header, err)
		return err
	}

	defer resp.Body.Close()

	return nil
}

// newFileUploadRequest create a file upload request
func newFileUploadRequest(url string, cpPath string, paths []string, params map[string]string) (*http.Request, error) {
	header := "migration.newFileUploadRequest"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for _, file := range paths {
		fullPath := cpPath + "/" + file
		logger := logrus.WithFields(logrus.Fields{
			"full path": fullPath,
		})
		logger.Debug("checkpoint files")

		fileP, err := os.Open(fullPath)
		if err != nil {
			logger.Error("open checkpoint file failed")
			return nil, err
		}

		part, err := writer.CreateFormFile(UploadFileKey, file)
		if err != nil {
			logger.Error("create checkpoint file failed")
			return nil, err
		}
		_, err = io.Copy(part, fileP)
		if err != nil {
			logger.Error("copy file failed")
			return nil, err
		}
		err = fileP.Close()
		if err != nil {
			logger.Error("close file failed")
			return nil, err
		}
	}

	// 其他参数列表写入 body
	for k, v := range params {
		if err := writer.WriteField(k, v); err != nil {
			logrus.WithFields(logrus.Fields{
				"key":   k,
				"value": v,
			}).Error("write param to request failed")
			return nil, err
		}
	}

	if err := writer.Close(); err != nil {
		logrus.Error("%s, close writer err: %v", header, err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		logrus.Error("%s, new http request err: %v", header, err)
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	return req, err
}

// getFilesFromCheckpoint get files from dir(pathname)
func getFilesFromCheckpoint(pathname string) ([]string, error) {
	header := "migration.getFilesFromCheckpoint"

	var files []string
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		logrus.Error("%s, read dir err: %v", header, err)
		return nil, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			// ignore files in dir criu.work/restore-...
			if restoreDir(fi.Name()) {
				continue
			}

			adder, err := getFilesFromCheckpoint(pathname + "/" + fi.Name())
			if err != nil {
				logrus.Error("%s, recursively getFilesFromCheckpoint err: %v", header, err)
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

// restoreDir check if path is dir or not
func restoreDir(filename string) bool {
	return strings.HasPrefix(filename, "restore-")
}
