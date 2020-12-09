package container

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
)

// StartOpts ...
type StartOpts struct {
	ContainerID string
	CStartOpts  types.ContainerStartOptions
}

// StartContainer start a container with opts
func StartContainer(opts StartOpts) error {
	header := "container.StartContainer"
	err := dockerCli.ContainerStart(context.Background(), opts.ContainerID, opts.CStartOpts)
	if err != nil {
		logrus.Errorf("%s, start container failed, err: %v", header, err)
		return err
	}
	return nil
}
