package client

import (
	"bytes"
	"github.com/yufeifly/migrator/api/types"

	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
)

// SendContainerCreate send container create request to dst node
func (cli *client) SendContainerCreate(options types.CreateReqOpts) ([]byte, error) {
	header := "client.SendContainerCreate"
	data := map[string]string{
		"ContainerName": options.ContainerName,
		"ImageName":     options.ImageName,
		"HostPort":      options.HostPort,
		"ContainerPort": options.ContainerPort,
		"PortBindings":  options.PortBindings,
		"ExposedPorts":  options.ExposedPorts,
		"Cmd":           options.Cmd,
	}

	ro := &grequests.RequestOptions{
		Data: data,
	}
	destUrl := "http://" + options.IP + ":" + options.Port + "/container/create"
	logrus.WithFields(logrus.Fields{
		"DestUrl": destUrl,
	}).Info(header)

	resp, err := grequests.Post(destUrl, ro)
	if err != nil {
		logrus.Errorf("%s, grequests.Post err: %v", header, err)
		return nil, err
	}
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.RawResponse.Body)
	if err != nil {
		logrus.Errorf("%s, read from response body err %v", header, err)
		return nil, err
	}
	defer resp.RawResponse.Body.Close()
	return body.Bytes(), nil
}
