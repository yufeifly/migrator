package logger

import (
	"github.com/yufeifly/migrator/api/types/log"
	"sync"
)

const capacity = 10

// Logger ...
type Logger struct {
	Count        int
	Capacity     int
	Sent         int
	Consumed     int
	sync.RWMutex // protect the log
	log.Log
	logBufferC chan log.Log
}

// NewLogger ...
func NewLogger() *Logger {
	return &Logger{
		Count:      0,
		Capacity:   capacity,
		Sent:       0,
		Consumed:   0,
		RWMutex:    sync.RWMutex{},
		Log:        log.Log{},
		logBufferC: make(chan log.Log, capacity),
	}
}

// ClearQueue clear data queue of the queue
func (l *Logger) ClearQueue() {
	l.LogQueue = l.LogQueue[:0]
}

// SetLastFlag set last flag of the log, which means logging will end
func (l *Logger) SetLastFlag() {
	l.Last = true
}

func (l *Logger) LogBuffer() chan log.Log {
	return l.logBufferC
}
