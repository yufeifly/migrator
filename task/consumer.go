package task

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types/logger"
	"github.com/yufeifly/migrator/client"
	"github.com/yufeifly/migrator/cluster"
	"github.com/yufeifly/migrator/redis"
	"time"
)

var DefaultConsumer *Consumer

type Consumer struct {
	ServicePort string
}

func NewConsumer() *Consumer {
	return &Consumer{
		ServicePort: "6379",
	}
}

// Consume consume a log in task queue
func (c *Consumer) Consume(ProxyServiceID, serviceID string) error {
	logrus.Debugf("Consume ProxyServiceID: %v", ProxyServiceID)
	proxy := cluster.DefaultCluster().GetProxy()
	cli := client.NewClient(proxy.Address)

	q := DefaultMapper.GetTaskQueue(ProxyServiceID)
	if q == nil {
		q := NewQueue()
		DefaultMapper.AddTaskQueue(ProxyServiceID, q)
		logrus.Warn("Consume: new a task queue")
	}

	// infinity loop, consume logs
	for {
		//fmt.Println("queue: ", DefaultQueue)
		logrus.Debug("tick")

		// get logs from the corresponding log queue
		taskJson := DefaultMapper.GetTaskQueue(ProxyServiceID).PopFront()
		// check if there are logs
		if taskJson == "" {
			time.Sleep(1 * time.Microsecond)
			continue
		}
		// unmarshall get serialized kv
		var task logger.Log
		err := json.Unmarshal([]byte(taskJson), &task)
		if err != nil {
			return err
		}
		if len(task.LogQueue) > 0 {
			for _, kv := range task.GetLogQueue() {
				var sli []string
				json.Unmarshal([]byte(kv), &sli)
				logrus.Infof("the slice: %v", sli)
				err := redis.Set(serviceID, sli[0], sli[1])
				if err != nil {
					logrus.Errorf("redis.set err: %v", err)
				}
			}
		}
		// stop this goroutine if it is the last task
		if task.GetLastFlag() {
			logrus.Warn("the last log consumed")
			return nil
		}
		logrus.Infof("consumed a log, send msg to src")
		err = cli.ConsumedAdder(ProxyServiceID)
		if err != nil {
			logrus.Errorf("cli.consumed failed, err: %v", err)
			return err
		}
	}
}
