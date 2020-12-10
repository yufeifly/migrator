package scheduler

import (
	"github.com/yufeifly/migrator/cusErr"
	"sync"
)

var DefaultScheduler *Scheduler

// InitScheduler ...
func InitScheduler() {
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

func (s *Scheduler) AddService(service *Service) {
	s.Map.Store(service.ID, service)
}

func (s *Scheduler) GetService(serviceID string) (*Service, error) {
	serviceP, ok := s.Map.Load(serviceID)
	if !ok {
		return nil, cusErr.ErrServiceNotFound
	}
	service, _ := serviceP.(*Service)
	return service, nil
}

func (s *Scheduler) DeleteService(serviceID string) {
	s.Map.Delete(serviceID)
}
