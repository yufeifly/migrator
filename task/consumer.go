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
	ServicePort string
}

func NewConsumer() *Consumer {
	return &Consumer{
		ServicePort: "6379",
	}
}

// Consume consume a log in task queue
// cid is the container id of src node
func (c *Consumer) Consume(cidDst, cidSrc string, node cluster.Node) error {

	cli := client.NewClient(node.Address)

	var t *Task
	// wait until log comes
	for {
		if t = Default().GetTask(cidSrc); t != nil {
			break
		}
		time.Sleep(time.Millisecond)
	}

	// consume logs
	for {
		logrus.Debug("tick")

		// get logs from the corresponding log queue
		select {
		case log := <-t.LogC:
			if len(log.GetLogQueue()) > 0 {
				for _, kv := range log.GetLogQueue() {
					err := redis.Set(cidDst, kv.Key, kv.Val)
					if err != nil {
						logrus.Errorf("redis.set err: %v", err)
					}
				}
			}
			// stop this goroutine if it is the last task
			if log.GetLastFlag() {
				logrus.Warn("the last log consumed")
				return nil
			}
			logrus.Infof("consumed a log, send msg to src")
			err := cli.ConsumedAdder(cidSrc)
			if err != nil {
				logrus.Errorf("cli.consumed failed, err: %v", err)
				return err
			}
		default:
			time.Sleep(1 * time.Millisecond)
		}
	}
}
