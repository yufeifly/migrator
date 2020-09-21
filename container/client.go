package container

import (
	"context"

	"github.com/docker/docker/client"
)

var (
	cli *client.Client
)

var ctx = context.Background()

func init() {
	var err error
	cli, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}
}
