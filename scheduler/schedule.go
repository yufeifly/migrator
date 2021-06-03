package scheduler

import (
	"github.com/yufeifly/migrator/cuserr"
	"sync"
)

var DefaultScheduler *Scheduler

// Init ...
func Init() {
	DefaultScheduler = NewScheduler()
}

type Scheduler struct {
	Map sync.Map
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

// Default get default scheduler
func Default() *Scheduler {
	return DefaultScheduler
}

func (s *Scheduler) AddContainerServ(containerServ *ContainerServ) {
	s.Map.Store(containerServ.CID, containerServ)
}

// GetContainerServ get container service by container id
func (s *Scheduler) GetContainerServ(cID string) (*ContainerServ, error) {
	cService, ok := s.Map.Load(cID)
	if !ok {
		return nil, cuserr.ErrContainerServiceNotFound
	}
	service, _ := cService.(*ContainerServ)
	return service, nil
}

func (s *Scheduler) DeleteService(serviceID string) {
	s.Map.Delete(serviceID)
}
