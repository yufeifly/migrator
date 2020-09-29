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
	ContainerName := c.PostForm("ContainerName")
	ImageName := c.PostForm("ImageName")
	HostPort := c.PostForm("HostPort")
	ContainerPort := c.PostForm("ContainerPort")
	PortBindings := c.PostForm("PortBindings")
	ExposedPorts := c.PostForm("ExposedPorts")
	CmdParam := c.PostForm("Cmd")

	// if HostPort & ContainerPort are set, then overwrite
	if HostPort != "" && ContainerPort != "" {
		openPort, _ := nat.NewPort("tcp", ContainerPort)
		exposedPorts := nat.PortSet{
			openPort: struct{}{}, //docker容器对外开放的端口
		}
		ExposedPortsSli, err := json.Marshal(exposedPorts)
		if err != nil {
			ReportErr(c, err)
			panic(err)
		}
		ExposedPorts = string(ExposedPortsSli)

		portBindings := nat.PortMap{
			openPort: []nat.PortBinding{{
				HostIP:   "0.0.0.0", //docker容器映射的宿主机的ip
				HostPort: HostPort,  //docker 容器映射到宿主机的端口
			}},
		}
		PortBindingsSli, err := json.Marshal(portBindings)
		if err != nil {
			ReportErr(c, err)
			panic(err)
		}
		PortBindings = string(PortBindingsSli)
	}

	createOpts := model.CreateOpts{
		ContainerName: ContainerName,
		ImageName:     ImageName,
		HostPort:      HostPort,
		ContainerPort: ContainerPort,
		ExposedPorts:  ExposedPorts,
		PortBindings:  PortBindings,
		Cmd:           CmdParam,
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
	// unmarshal cmd
	var cmd []string
	err := json.Unmarshal([]byte(opts.Cmd), &cmd)
	if err != nil {
		fmt.Printf("unmarshal err: %v\n", err)
		return container.ContainerCreateCreatedBody{}, err
	}

	config := &container.Config{
		Image: opts.ImageName,
		Cmd:   cmd,
	}

	if opts.ExposedPorts != "" {
		exposedPorts := nat.PortSet{}
		err = json.Unmarshal([]byte(opts.ExposedPorts), &exposedPorts)
		if err != nil {
			fmt.Printf("unmarshal err: %v\n", err)
			return container.ContainerCreateCreatedBody{}, err
		}
		config.ExposedPorts = exposedPorts
	}

	hostConfig := &container.HostConfig{}
	if opts.PortBindings != "" {
		portBindings := nat.PortMap{}
		err = json.Unmarshal([]byte(opts.PortBindings), &portBindings)
		if err != nil {
			fmt.Printf("unmarshal err: %v\n", err)
			return container.ContainerCreateCreatedBody{}, err
		}
		hostConfig.PortBindings = portBindings
	}

	body, err := cli.ContainerCreate(ctx, config, hostConfig, nil, opts.ContainerName)
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}

	fmt.Printf("Create container ID: %s\n", body.ID)

	return body, err
}
