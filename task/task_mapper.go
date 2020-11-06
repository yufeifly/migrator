// map[ProxyServiceID]taskQueue
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

func (m *Mapper) AddTaskQueue(ProxyServiceID string, q *Queue) {
	m.Lock()
	if m.GetTaskQueue(ProxyServiceID) == nil {
		m.Map.Store(ProxyServiceID, q)
	}
	m.Unlock()
}

// GetTaskQueue get task queue for a ProxyService
func (m *Mapper) GetTaskQueue(ProxyServiceID string) *Queue {
	que, ok := m.Map.Load(ProxyServiceID)
	if !ok {
		return nil
	}
	q, _ := que.(*Queue)
	return q
}
