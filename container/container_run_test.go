package container

import (
	"fmt"
	"github.com/yufeifly/migrator/model"
	"testing"
)

func TestRunContainer(t *testing.T) {
	_, err := RunContainer(model.RunOpts{
		ImageName:     "docker.io/library/redis",
		HostPort:      "8998",
		ContainerPort: "8998",
	},
	)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("pass")
	}
}
