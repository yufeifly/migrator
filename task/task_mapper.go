// map[ProxyServiceID]taskQueue
package task

import (
	"github.com/yufeifly/migrator/api/types/log"
	"sync"
)

type Task struct {
	CID  string
	LogC chan log.Log
}

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

func Default() *Mapper {
	return DefaultMapper
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

func (m *Mapper) GetTask(cid string) *Task {
	task, ok := m.Map.Load(cid)
	if !ok {
		return nil
	}
	t, _ := task.(*Task)
	return t
}
