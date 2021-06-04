// map[ProxyServiceID]taskQueue
package task

import (
	"sync"
)

var defaultMapper *Mapper

func init() {
	defaultMapper = NewMapper()
}

type Mapper struct {
	sync.Map
	sync.Mutex
}

func NewMapper() *Mapper {
	return &Mapper{}
}

func Default() *Mapper {
	return defaultMapper
}

func (m *Mapper) AddTask(cid string, t *Task) {
	m.Lock()
	if m.GetTask(cid) == nil {
		m.Map.Store(cid, t)
	}
	m.Unlock()
}

func (m *Mapper) GetTask(cid string) *Task {
	task, ok := m.Map.Load(cid)
	if !ok {
		return nil
	}
	t, _ := task.(*Task)
	return t
}
