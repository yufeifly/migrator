package container

import (
	"fmt"
	"github.com/yufeifly/proxyd/model"
	"testing"
)

func TestCreateContainer(t *testing.T) {
	createOpts := model.CreateOpts{
		ContainerName: "",
		ImageName:     "",
		HostPort:      "",
		ContainerPort: "",
		Cmd:           "",
		DestIP:        "",
	}
	_, err := CreateContainer(createOpts)
	if err != nil {
		fmt.Println("TestCreateContainer err: ", err)
	}
}
