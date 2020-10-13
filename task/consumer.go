package task

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/model"
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
func (c *Consumer) Consume() error {
	// infinity loop, consume logs
	for {
		//fmt.Println("queue: ", DefaultQueue)
		taskJson := DefaultQueue.PopFront()
		if taskJson == "" {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		fmt.Printf("taskJson: %v\n", taskJson)
		// unmarshall get serialized kv
		var task model.Log
		err := json.Unmarshal([]byte(taskJson), &task)
		if err != nil {
			return err
		}
		fmt.Printf("consumer log: %v\n", task)
		for _, kv := range task.GetLogQueue() {
			var sli []string
			json.Unmarshal([]byte(kv), &sli)
			redis.Set(sli[0], sli[1])
		}
		// consumed a log, send this message to src
		logrus.Infof("consumed a log, msg send to src")
		time.Sleep(200 * time.Millisecond)
		// stop this goroutine if it is the last task
		flag := task.GetLastFlag()
		fmt.Println("flag: ", flag)
		if flag {
			return nil
		}
	}
	return nil
}
