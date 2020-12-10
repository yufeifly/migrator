package logger

type Log struct {
	Last     bool
	LogQueue []string
}

type LogWithServiceID struct {
	Log
	ProxyServiceID string
}

func NewLog() *Log {
	return &Log{
		Last:     false,
		LogQueue: []string{},
	}
}

func (l *Log) GetLastFlag() bool {
	return l.Last
}

func (l *Log) GetLogQueue() []string {
	return l.LogQueue
}

func (l *Log) SetLastFlag(f bool) {
	l.Last = f
}

func (l *Log) SetLogQueue(s []string) {
	l.LogQueue = s
}
