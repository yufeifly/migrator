// map[serviceID]taskQueue
package task

import "sync"

var DefaultMapper *Mapper

func init() {
	DefaultMapper = NewMapper()
}

type Mapper struct {
	sync.Map
	sync.Mutex
}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (m *Mapper) AddTaskQueue(serviceID string, q *Queue) {
	m.Lock()
	if m.GetTaskQueue(serviceID) == nil {
		m.Map.Store(serviceID, q)
	}
	m.Unlock()
}

func (m *Mapper) GetTaskQueue(serviceID string) *Queue {
	que, ok := m.Map.Load(serviceID)
	if !ok {
		return nil
	}
	q, _ := que.(*Queue)
	return q
}
