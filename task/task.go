package task

import "github.com/yufeifly/migrator/api/types/log"

const capacity = 100

type Task struct {
	CID  string
	LogC chan log.Log
}

func NewTask(cid string) *Task {
	return &Task{
		CID:  cid,
		LogC: make(chan log.Log, capacity),
	}
}

func (t *Task) Push(log log.Log) {
	t.LogC <- log
}

func (t *Task) Get() log.Log {
	log := <-t.LogC
	return log
}
