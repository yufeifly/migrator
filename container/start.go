package container

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/model"
)

// StartContainer start a container with opts
func StartContainer(opts model.StartOpts) error {
	header := "container.StartContainer"
	err := cli.ContainerStart(context.Background(), opts.ContainerID, opts.CStartOpts)
	if err != nil {
		logrus.Errorf("%s, start container failed, err: %v", header, err)
		return err
	}
	return nil
}
