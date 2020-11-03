package container

import (
	"github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/model"
)

// RunContainer
func RunContainer(opts model.RunOpts) (string, error) {
	header := "container.RunContainer"
	out, err := cli.ImagePull(ctx, opts.ImageName, types.ImagePullOptions{})
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Infof("%s, pull image result: %v", header, out)

	resp, err := CreateContainer(model.CreateOpts{
		ImageName:     opts.ImageName,
		HostPort:      opts.HostPort,
		ContainerPort: opts.ContainerPort,
	})
	if err != nil {
		logrus.Panic(err)
	}

	err = StartContainer(model.StartOpts{
		ContainerID: resp.ID,
		CStartOpts:  types.ContainerStartOptions{},
	})
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Infof("container id: %v", resp.ID)
	return resp.ID, nil
}
