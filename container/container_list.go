package container

import (
	"github.com/docker/docker/api/types"
)

// ListContainers
func ListContainers(opts types.ContainerListOptions) ([]types.Container, error) {
	containers, err := cli.ContainerList(ctx, opts)
	if err != nil {
		return nil, err
	}
	return containers, nil
}
