package client

import (
	"encoding/json"
	"fmt"
	"github.com/yufeifly/proxyd/model"
	"testing"
)

func TestCli_SendContainerCreate(t *testing.T) {
	c := Cli{}
	cmdSlice := []string{"/bin/sh", "-c", "i=0; while true; do echo $i; i=$(expr $i + 1); sleep 1; done"}
	cmd, err := json.Marshal(&cmdSlice)
	fmt.Printf("cmd: %v\n", string(cmd))

	opts := model.CreateOpts{
		ContainerName: "bb21",
		ImageName:     "busybox",
		HostPort:      "",
		ContainerPort: "",
		Cmd:           string(cmd),
		DestIP:        "http://127.0.0.1:6789/docker/create",
	}
	got, err := c.SendContainerCreate(opts)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Printf("create result: %v\n", string(got))
	}
}
