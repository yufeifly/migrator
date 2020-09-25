package model

import "github.com/docker/docker/api/types"

type StartOpts struct {
	CStartOpts  types.ContainerStartOptions
	ContainerID string
}
