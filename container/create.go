package container

import (
	"encoding/json"
	"github.com/docker/go-connections/nat"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxyd/model"
	"github.com/yufeifly/proxyd/utils"

	"github.com/docker/docker/api/types/container"
	"github.com/gin-gonic/gin"
)

// Create handler for creating a container
func Create(c *gin.Context) {
	header := "container.Create"

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
			utils.ReportErr(c, err)
			logrus.Panic(err)
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
			utils.ReportErr(c, err)
			logrus.Panic(err)
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
		utils.ReportErr(c, err)
		logrus.Errorf("%s, CreateContainer err: %v", header, err)
		logrus.Panic(err)
	}
	c.JSON(200, gin.H{
		"result":      "success",
		"containerId": body.ID,
	})
}

// CreateContainer create a container
func CreateContainer(opts model.CreateOpts) (container.ContainerCreateCreatedBody, error) {
	header := "container.CreateContainer"

	config := &container.Config{
		Image: opts.ImageName,
	}
	// unmarshal cmd
	if opts.Cmd != "" {
		var cmd []string
		err := json.Unmarshal([]byte(opts.Cmd), &cmd)
		if err != nil {
			logrus.Errorf("%s, unmarshal cmd err: %v", header, err)
			return container.ContainerCreateCreatedBody{}, err
		}
		config.Cmd = cmd
	}

	if opts.ExposedPorts != "" {
		exposedPorts := nat.PortSet{}
		err := json.Unmarshal([]byte(opts.ExposedPorts), &exposedPorts)
		if err != nil {
			logrus.Errorf("%s, unmarshal ExposedPorts err: %v", header, err)
			return container.ContainerCreateCreatedBody{}, err
		}
		config.ExposedPorts = exposedPorts
	}

	hostConfig := &container.HostConfig{}
	if opts.PortBindings != "" {
		portBindings := nat.PortMap{}
		err := json.Unmarshal([]byte(opts.PortBindings), &portBindings)
		if err != nil {
			logrus.Errorf("%s, unmarshal PortBindings err: %v", header, err)
			return container.ContainerCreateCreatedBody{}, err
		}
		hostConfig.PortBindings = portBindings
	}

	body, err := cli.ContainerCreate(ctx, config, hostConfig, nil, opts.ContainerName)
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}

	logrus.WithFields(logrus.Fields{
		"ContainerID": body.ID,
	}).Info("container created")

	return body, err
}
