package task

import (
	"sync"
)

type Queue struct {
	Q        []string
	TaskLeft int // current task number
	TotalNum int // task consumed
	sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		Q:        []string{},
		TaskLeft: 0,
		TotalNum: 0,
		Mutex:    sync.Mutex{},
	}
}

func (q *Queue) GetTaskLeft() int {
	var ret int
	q.Lock()
	ret = q.TaskLeft
	q.Unlock()
	return ret
}

func (q *Queue) GetTotalTask() int {
	var ret int
	q.Lock()
	ret = q.TotalNum
	q.Unlock()
	return ret
}

// Push push a serialized task
func (q *Queue) Push(task string) {
	q.Lock()
	q.TaskLeft++
	q.Q = append(q.Q, task)
	q.Unlock()
}

// PopFront get a task from queue, if there is no task, return ""
func (q *Queue) PopFront() string {
	var task string
	q.Lock()
	if q.TaskLeft == 0 {
		task = ""
	} else {
		q.TotalNum++
		q.TaskLeft--
		task = q.Q[0]
		q.Q = q.Q[1:]
	}
	q.Unlock()
	return task
}
