package client

import (
	"bytes"

	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/model"
)

// SendContainerCreate send container create request to dst node
func (cli *Client) SendContainerCreate(opts model.CreateReqOpts) ([]byte, error) {
	header := "client.SendContainerCreate"
	data := map[string]string{
		"ContainerName": opts.ContainerName,
		"ImageName":     opts.ImageName,
		"HostPort":      opts.HostPort,
		"ContainerPort": opts.ContainerPort,
		"PortBindings":  opts.PortBindings,
		"ExposedPorts":  opts.ExposedPorts,
		"Cmd":           opts.Cmd,
	}

	ro := &grequests.RequestOptions{
		Data: data,
	}
	destUrl := "http://" + opts.IP + ":" + opts.Port + "/container/create"
	logrus.WithFields(logrus.Fields{
		"DestUrl": destUrl,
	}).Info(header)

	resp, err := grequests.Post(destUrl, ro)
	if err != nil {
		return nil, err
	}
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.RawResponse.Body)
	if err != nil {
		logrus.Errorf("%s: read from response body err %v", header, err)
		return nil, err
	}
	defer resp.RawResponse.Body.Close()
	return body.Bytes(), nil
}
