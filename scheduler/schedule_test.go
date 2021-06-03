package scheduler

import (
	"fmt"
	"testing"
)

func TestScheduler_AddService(t *testing.T) {
	service := &ContainerServ{
		SID:        "service1",
		CID:        "123456",
		ServiceCli: nil,
	}
	DefaultScheduler.AddContainerServ(service)
	get, _ := DefaultScheduler.GetContainerServ(service.CID)
	fmt.Printf("get serive: %v\n", get)
}
