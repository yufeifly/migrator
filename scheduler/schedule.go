package scheduler

import (
	"github.com/yufeifly/migrator/cusErr"
	"sync"
)

var DefaultScheduler *Scheduler

func init() {
	DefaultScheduler = NewScheduler()
}

type Scheduler struct {
	Map sync.Map
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
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
