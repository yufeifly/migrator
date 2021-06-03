package log

type KV struct {
	Key string
	Val string
}

type Log struct {
	Last     bool
	LogQueue []KV
}

func NewLog() *Log {
	return &Log{
		Last:     false,
		LogQueue: []KV{},
	}
}

func (l *Log) GetLastFlag() bool {
	return l.Last
}

func (l *Log) GetLogQueue() []KV {
	return l.LogQueue
}

func (l *Log) SetLastFlag(f bool) {
	l.Last = f
}

func (l *Log) SetLogQueue(s []KV) {
	l.LogQueue = s
}

// LogWithCID ...
type LogWithCID struct {
	CID string
	Log
}

// NewLogWithCID new a log with container id of the target node
func NewLogWithCID(cid string) *LogWithCID {
	return &LogWithCID{
		CID: cid,
		Log: Log{},
	}
}
