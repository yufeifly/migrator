package container

import (
	"fmt"
	"github.com/docker/docker/api/types"
)

// Inspect get info of container
func Inspect(containerID string) (types.ContainerJSON, error) {
	serverResp, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return types.ContainerJSON{}, err
	}
	return serverResp, err
}

// @deprecated get image name
func GetImageByImageID(imageID string) (string, error) {
	imageInspect, _, err := cli.ImageInspectWithRaw(ctx, imageID)
	if err != nil {
		return "", err
	}
	fmt.Printf("image: %v\n", imageInspect)
	return imageInspect.RepoTags[0], err
}

// GetImageByContainer get image name of container, not directly used
func GetImageByContainer(containerID string) (string, error) {
	serverResp, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return "", err
	}
	return serverResp.Config.Image, err
}

//
func GetPortBinding(containerID string) error {
	Resp, err := Inspect(containerID)
	if err != nil {
		return err
	}
	fmt.Printf("portBinding: %v\n", Resp.HostConfig.PortBindings)
	return nil
}
