package container

import (
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

func Inspect(containerID string) (types.ContainerJSON, error) {
	serverResp, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return types.ContainerJSON{}, err
	}
	return serverResp, err
}

func GetImageRepoTags(containerName string) (string, error) {
	//imageInspect, _, err := cli.ImageInspectWithRaw(ctx, imageID)
	//if err != nil {
	//	return nil, err
	//}
	//return imageInspect.RepoTags, err
	p := "name=" + containerName
	fmt.Printf("filter: %v\n", p)
	filter, err := filters.FromParam(p)
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Filters: filter,
	})
	if err != nil {
		return "", err
	}
	if len(containers) <= 0 {
		return "", errors.New("no such container")
	}
	return containers[0].Image, nil
}

func GetImageByImageID(imageID string) (string, error) {
	imageInspect, _, err := cli.ImageInspectWithRaw(ctx, imageID)
	if err != nil {
		return "", err
	}
	fmt.Printf("image: %v\n", imageInspect)
	return imageInspect.RepoTags[0], err
}
