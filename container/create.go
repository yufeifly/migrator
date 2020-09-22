package container

import (
	"fmt"
	"github.com/docker/go-connections/nat"

	"github.com/docker/docker/api/types/container"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	containerName := c.Request.URL.Query().Get("containerName")
	imageName := c.Request.URL.Query().Get("imageName")
	hostPort := c.Request.URL.Query().Get("hostPort")

	exports := make(nat.PortSet, 10)
	port, err := nat.NewPort("tcp", "80")
	if err != nil {
		ReportErr(c, err)
		panic(err)
	}

	exports[port] = struct{}{}
	config := &container.Config{
		Image:        imageName,
		ExposedPorts: exports,
	}

	portBind := nat.PortBinding{
		HostPort: hostPort,
	}
	portMap := make(nat.PortMap, 0)
	tmp := make([]nat.PortBinding, 0, 1)
	tmp = append(tmp, portBind)
	portMap[port] = tmp
	hostConfig := &container.HostConfig{PortBindings: portMap}

	body, err := cli.ContainerCreate(ctx, config, hostConfig, nil, containerName)
	if err != nil {
		ReportErr(c, err)
		panic(err)
	}

	fmt.Printf("Create container ID: %s\n", body.ID)

	c.JSON(200, gin.H{
		"result":      "success",
		"containerId": body.ID,
	})
}
