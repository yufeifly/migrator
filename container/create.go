package container

import (
	"encoding/json"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/yufeifly/proxyd/model"

	"github.com/docker/docker/api/types/container"
	"github.com/gin-gonic/gin"
)

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

	var cmd []string
	err := json.Unmarshal([]byte(cmdParam), &cmd)
	if err != nil {
		fmt.Printf("unmarshal err: %v\n", err)
		ReportErr(c, err)
		panic(err)
	}

	config := &container.Config{
		Image: imageName,
		Cmd:   cmd,
	}

	hostConfig := &container.HostConfig{}

	if hostPort != "" && containerPort != "" {
		openPort, _ := nat.NewPort("tcp", containerPort)
		config.ExposedPorts = nat.PortSet{
			openPort: struct{}{}, //docker容器对外开放的端口
		}

		hostConfig.PortBindings = nat.PortMap{
			openPort: []nat.PortBinding{nat.PortBinding{
				HostIP:   "0.0.0.0", //docker容器映射的宿主机的ip
				HostPort: hostPort,  //docker 容器映射到宿主机的端口
			}},
		}
	}

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

func CreateContainer(opts model.CreateOpts) error {

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

	return err
}
