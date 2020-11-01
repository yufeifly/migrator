package model

import "github.com/docker/docker/api/types"

type StartOpts struct {
	ContainerID string
	CStartOpts  types.ContainerStartOptions
}
