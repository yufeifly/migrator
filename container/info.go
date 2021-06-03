package container

import (
	"github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
)

// Inspect get info of container
func Inspect(containerID string) (types.ContainerJSON, error) {
	resp, err := dockerCli.ContainerInspect(ctx, containerID)
	if err != nil {
		return types.ContainerJSON{}, err
	}
	return resp, err
}

// @deprecated replaced by GetImageByContainer
// GetImageByImageID get image name
func GetImageByImageID(imageID string) (string, error) {
	imageInspect, _, err := dockerCli.ImageInspectWithRaw(ctx, imageID)
	if err != nil {
		return "", err
	}
	logrus.Infof("image: %v", imageInspect)
	return imageInspect.RepoTags[0], err
}

// GetImageByContainer get image name of container, not directly used
func GetImageByContainer(containerID string) (string, error) {
	serverResp, err := dockerCli.ContainerInspect(ctx, containerID)
	if err != nil {
		return "", err
	}
	return serverResp.Config.Image, err
}

// GetPortBinding get port binding of container
func GetPortBinding(containerID string) error {
	Resp, err := Inspect(containerID)
	if err != nil {
		return err
	}
	logrus.Infof("portBinding: %v", Resp.HostConfig.PortBindings)
	return nil
}
