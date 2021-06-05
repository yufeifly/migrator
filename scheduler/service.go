package scheduler

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types"
	"github.com/yufeifly/migrator/api/types/log"
	"github.com/yufeifly/migrator/api/types/logger"
	"github.com/yufeifly/migrator/api/types/svc"
	"github.com/yufeifly/migrator/client"
	"github.com/yufeifly/migrator/ticket"
)

type ContainerServ struct {
	CID        string         // container id
	SID        string         // service id
	Port       string         // service port, also the exposed port of the container
	logger     *logger.Logger // log the data while migrating, useful in migration
	ticket     ticket.Ticket  // ticket interface
	ServiceCli *redis.Client  // redis connection
}

// NewService new a storage service, keep it in map
func NewContainerServ(opts svc.ServiceOpts) *ContainerServ {
	return &ContainerServ{
		CID:    opts.CID,
		SID:    opts.SID,
		Port:   opts.Port,
		logger: logger.NewLogger(),
		ticket: ticket.NewTicket(),
		ServiceCli: redis.NewClient(&redis.Options{
			Addr:     "localhost" + ":" + opts.Port,
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

// RegisterServices register services
func RegisterServices() {
	opts := []svc.ServiceOpts{
		{
			CID:  "s1.c1",
			SID:  "s1",
			Port: "39955",
		},
		{
			CID:  "s1.c2",
			SID:  "s1",
			Port: "39956",
		},
	}
	for _, opt := range opts {
		DefaultScheduler.AddContainerServ(NewContainerServ(opt))
	}
}

// LockAndGetSentConsumed return sent and consumed
func (s *ContainerServ) LockAndGetSentConsumed() (int, int) {
	s.logger.Lock()
	defer s.logger.Unlock()
	return s.logger.Sent, s.logger.Consumed
}

// LoggingFinished check if the logging process finished
func (s *ContainerServ) LoggingFinished() bool {
	s.logger.RLock()
	defer s.logger.RUnlock()
	sent, consumed := s.logger.Sent, s.logger.Consumed
	if sent > 0 && sent-consumed < 1 {
		return true
	}
	return false
}

// LogSent get the number of logs already sent
func (s *ContainerServ) LogSent() int {
	s.logger.RLock()
	defer s.logger.RUnlock()
	return s.logger.Sent
}

func (s *ContainerServ) SendLog(l log.Log, target types.Address, last bool) error {
	cli := client.NewClient(target)
	if last {
		l.SetLastFlag(true)
	}
	logWithCID := log.LogWithCID{
		Log: l,
		CID: s.CID,
	}
	err := cli.SendLog(logWithCID)
	if err != nil {
		return err
	}
	s.logger.Lock()
	s.logger.Sent++
	s.logger.Unlock()
	return nil
}

// ConsumedAdder ...
func (s *ContainerServ) ConsumedAdder() {
	s.logger.Lock()
	s.logger.Consumed++
	s.logger.Unlock()
}

// Ticket get ticket interface
func (s *ContainerServ) Ticket() ticket.Ticket {
	return s.ticket
}

func (s *ContainerServ) LogRecord(key, val string) error {
	logrus.Warn("redis.logRecord logging operation")
	kv := log.KV{
		Key: key,
		Val: val,
	}
	s.logger.Lock()
	defer s.logger.Unlock()
	s.logger.Count++
	s.logger.LogQueue = append(s.logger.LogQueue, kv)

	if s.logger.Count == s.logger.Capacity {
		s.logger.LogBuffer() <- s.logger.Log
		s.logger.ClearQueue()
		s.logger.Count = 0
	}
	return nil
}

func (s *ContainerServ) Logger() *logger.Logger {
	return s.logger
}
