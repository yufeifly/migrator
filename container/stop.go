package container

import (
	"time"
)

func StopContainer(ContainerID string, timeout time.Duration) error {
	err := cli.ContainerStop(ctx, ContainerID, &timeout)
	if err != nil {
		return err
	}
	return nil
}
