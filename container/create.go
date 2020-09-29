package container

import (
	"encoding/json"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/yufeifly/proxyd/model"

	"github.com/docker/docker/api/types/container"
	"github.com/gin-gonic/gin"
)

// Create handler for creating a container
func Create(c *gin.Context) {
	//containerName := c.Request.URL.Query().Get("containerName")
	//imageName := c.Request.URL.Query().Get("imageName")
	//hostPort := c.Request.URL.Query().Get("hostPort")
	//containerPort := c.Request.URL.Query().Get("containerPort")
	containerName := c.PostForm("containerName")
	imageName := c.PostForm("imageName")
	hostPort := c.PostForm("hostPort")
	containerPort := c.PostForm("containerPort")
	cmdParam := c.PostForm("cmd")

	//fmt.Printf("containerName: %v\n", containerName)
	//fmt.Printf("cmdParam: %v\n", cmdParam)

	createOpts := model.CreateOpts{
		ContainerName: containerName,
		ImageName:     imageName,
		HostPort:      hostPort,
		ContainerPort: containerPort,
		Cmd:           cmdParam,
	}
	body, err := CreateContainer(createOpts)
	if err != nil {
		ReportErr(c, err)
		fmt.Printf("CreateContainer err: %v\n", err)
		panic(err)
	}
	c.JSON(200, gin.H{
		"result":      "success",
		"containerId": body.ID,
	})
}

// CreateContainer create a container
func CreateContainer(opts model.CreateOpts) (container.ContainerCreateCreatedBody, error) {

	var cmd []string
	err := json.Unmarshal([]byte(opts.Cmd), &cmd)
	if err != nil {
		fmt.Printf("unmarshal err: %v\n", err)
		panic(err)
	}

	config := &container.Config{
		Image: opts.ImageName,
		Cmd:   cmd,
	}

	hostConfig := &container.HostConfig{}

	if opts.HostPort != "" && opts.ContainerPort != "" {
		openPort, _ := nat.NewPort("tcp", opts.ContainerPort)
		config.ExposedPorts = nat.PortSet{
			openPort: struct{}{}, //docker容器对外开放的端口
		}

		hostConfig.PortBindings = nat.PortMap{
			openPort: []nat.PortBinding{nat.PortBinding{
				HostIP:   "0.0.0.0",     //docker容器映射的宿主机的ip
				HostPort: opts.HostPort, //docker 容器映射到宿主机的端口
			}},
		}
	}

	body, err := cli.ContainerCreate(ctx, config, hostConfig, nil, opts.ContainerName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Create container ID: %s\n", body.ID)

	return body, err
}
