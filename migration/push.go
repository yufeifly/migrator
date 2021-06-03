package migration

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

var logger = logrus.New()

// PushOpts push checkpoint to dst node
type PushOpts struct {
	CheckPointID  string
	CheckPointDir string
	Dest          types.Address
	CID           string
	SID           string
	Port          string
}

// PushCheckpoint push checkpoint to destination and deliver restore request
func PushCheckpoint(migOpts PushOpts) error {
	header := "migration.PushCheckpoint"

	ip := migOpts.Dest.IP
	port := migOpts.Dest.Port

	urlPost := "http://" + ip + ":" + port + "/container/checkpoint/restore"
	params := map[string]string{
		"ContainerID":   migOpts.CID,
		"ServiceID":     migOpts.SID,
		"ExposedPort":   migOpts.Port,
		"CheckPointID":  migOpts.CheckPointID,
		"CheckPointDir": migOpts.CheckPointDir,
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
		logEntry := logger.WithFields(logrus.Fields{
			"full path": fullPath,
		})
		logEntry.Debug("checkpoint files")

		fileP, err := os.Open(fullPath)
		if err != nil {
			logEntry.Error("open checkpoint file failed")
			return nil, err
		}

		part, err := writer.CreateFormFile(UploadFileKey, file)
		if err != nil {
			logEntry.Error("create checkpoint file failed")
			return nil, err
		}
		_, err = io.Copy(part, fileP)
		if err != nil {
			logEntry.Error("copy file failed")
			return nil, err
		}
		err = fileP.Close()
		if err != nil {
			logEntry.Error("close file failed")
			return nil, err
		}
	}

	// other params write to body
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
		logrus.Errorf("%s, close writer err: %v", header, err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		logrus.Errorf("%s, new http request err: %v", header, err)
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
		logrus.Errorf("%s, read dir err: %v", header, err)
		return nil, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			// ignore files in dir criu.work, e.g. restore-...
			if restoreDir(fi.Name()) {
				continue
			}

			adder, err := getFilesFromCheckpoint(pathname + "/" + fi.Name())
			if err != nil {
				logrus.Errorf("%s, recursively getFilesFromCheckpoint err: %v", header, err)
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
