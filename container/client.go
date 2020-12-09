package container

import (
	"context"
	"github.com/sirupsen/logrus"

	"github.com/docker/docker/client"
)

var dockerCli *client.Client

var ctx = context.Background()

func init() {
	var err error
	dockerCli, err = client.NewEnvClient()
	if err != nil {
		logrus.Panic(err)
	}
}
