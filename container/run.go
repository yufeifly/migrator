package container

import (
	dockertypes "github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types"
)

// RunOpts ...
type RunOpts struct {
	ImageName     string
	HostPort      string
	ContainerPort string
}

// RunContainer run a container
func RunContainer(opts RunOpts) (string, error) {
	header := "container.RunContainer"
	out, err := dockerCli.ImagePull(ctx, opts.ImageName, dockertypes.ImagePullOptions{})
	if err != nil {
		logrus.Panicf("%s, dockerCli.ImagePull err: %v", header, err)
	}
	logrus.Infof("%s, pull image result: %v", header, out)

	resp, err := CreateContainer(types.CreateOpts{
		ImageName:     opts.ImageName,
		HostPort:      opts.HostPort,
		ContainerPort: opts.ContainerPort,
	})
	if err != nil {
		logrus.Panicf("%s, CreateContainer err: %v", header, err)
	}

	err = StartContainer(StartOpts{
		ContainerID: resp.ID,
		CStartOpts:  dockertypes.ContainerStartOptions{},
	})
	if err != nil {
		logrus.Panicf("%s, StartContainer err: %v", header, err)
	}
	logrus.Infof("%s, container id: %v", header, resp.ID)
	return resp.ID, nil
}
