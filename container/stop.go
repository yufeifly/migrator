package container

import (
	"github.com/sirupsen/logrus"
	"time"
)

// StopContainer stop the container
func StopContainer(ContainerID string, timeout *time.Duration) error {
	header := "container.StopContainer"
	err := cli.ContainerStop(ctx, ContainerID, timeout)
	if err != nil {
		logrus.Errorf("%s, start container failed, err: %v", header, err)
		return err
	}
	return nil
}
