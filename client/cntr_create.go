package client

import (
	"bytes"
	"fmt"
	"github.com/yufeifly/proxyd/model"
	"mime/multipart"
	"net/http"
)

func (c *Cli) SendContainerCreate(opts model.CreateOpts) ([]byte, error) {
	params := map[string]string{
		"containerName": opts.ContainerName,
		"imageName":     opts.ImageName,
		"hostPort":      opts.HostPort,
		"containerPort": opts.ContainerPort,
		"cmd":           opts.Cmd,
	}

	req, err := NewCreateRequest(opts.DestIP, params)
	if err != nil {
		fmt.Printf("error to new upload file request:%s\n", err.Error())
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error to request to the server:%s\n", err.Error())
		return nil, err
	}

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		fmt.Printf("error to request to the server:%s\n", err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	return body.Bytes(), nil
}

func NewCreateRequest(url string, params map[string]string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
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
