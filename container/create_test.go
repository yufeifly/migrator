package container

import (
	"fmt"
	"github.com/yufeifly/migrator/api/types"
	"testing"
)

func TestCreateContainer(t *testing.T) {
	createOpts := types.CreateOpts{
		ContainerName: "",
		ImageName:     "",
		HostPort:      "",
		ContainerPort: "",
		Cmd:           "",
	}
	_, err := CreateContainer(createOpts)
	if err != nil {
		fmt.Println("TestCreateContainer err: ", err)
	}
}
