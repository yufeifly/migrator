package task

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/client"
	"github.com/yufeifly/migrator/cluster"
	"github.com/yufeifly/migrator/redis"
	"time"
)

var DefaultConsumer *Consumer

type Consumer struct {
}

func NewConsumer() *Consumer {
	return &Consumer{}
}

// Consume consume a log in task queue
func (c *Consumer) Consume(cid string, node cluster.Node) error {

	cli := client.NewClient(node.Address)

	var t *Task
	// wait until log comes
	for {
		if t = Default().GetTask(cid); t != nil {
			break
		}
		time.Sleep(time.Millisecond)
	}
	// consume logs
	for {
		select {
		case log := <-t.LogC:
			if len(log.GetLogQueue()) > 0 {
				for _, kv := range log.GetLogQueue() {
					err := redis.Set(cid, kv.Key, kv.Val)
					if err != nil {
						logrus.Errorf("redis.set err: %v", err)
					}
				}
			}
			logrus.Infof("consumed a log, send msg to src")
			err := cli.ConsumedAdder(cid)
			if err != nil {
				logrus.Errorf("cli.consumed failed, err: %v", err)
				return err
			}
			// stop this goroutine if it is the last task
			if log.GetLastFlag() {
				logrus.Warn("the last log consumed")
				return nil
			}
		default:
			time.Sleep(1 * time.Millisecond)
		}
	}
}
