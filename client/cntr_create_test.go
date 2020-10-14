package client

import (
	"encoding/json"
	"fmt"
	"github.com/yufeifly/migrator/model"
	"testing"
)

func TestCli_SendContainerCreate(t *testing.T) {
	c := Client{}
	cmdSlice := []string{"/bin/sh", "-c", "i=0; while true; do echo $i; i=$(expr $i + 1); sleep 1; done"}
	cmd, err := json.Marshal(&cmdSlice)
	fmt.Printf("cmd: %v\n", string(cmd))

	opts := model.CreateReqOpts{
		CreateOpts: model.CreateOpts{
			ContainerName: "bb22",
			ImageName:     "busybox",
			HostPort:      "",
			ContainerPort: "",
			Cmd:           string(cmd),
		},
		DestIP:   "127.0.0.1",
		DestPort: "6789",
	}
	got, err := c.SendContainerCreate(opts)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		var ans map[string]interface{}
		json.Unmarshal(got, &ans)
		fmt.Printf("create result: %v\n", ans["containerId"])
	}
}
