package scheduler

import (
	"fmt"
	"testing"
)

func TestScheduler_AddService(t *testing.T) {
	service := &Service{
		ID:          "service1",
		ContainerID: "123456",
		ServiceCli:  nil,
	}
	DefaultScheduler.AddService(service)
	get, _ := DefaultScheduler.GetService(service.ID)
	fmt.Printf("get serive: %v\n", get)
}
