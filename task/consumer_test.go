package task

import (
	"encoding/json"
	"fmt"
	"github.com/yufeifly/migrator/model"
	"strconv"
	"testing"
)

func TestConsumer_Consume(t *testing.T) {
	log := model.NewLog()
	var data []string
	for i := 0; i < 50; i++ {
		s := strconv.Itoa(i)
		tmp := []string{s, s + "#"}
		tmpJson, _ := json.Marshal(tmp)
		data = append(data, string(tmpJson))
	}
	log.SetLogQueue(data)
	f := true
	log.SetLastFlag(f)

	fmt.Printf("log: %v\n", log)

	logJson, _ := json.Marshal(log)
	fmt.Println("logjson: ", string(logJson))
	DefaultQueue.Push(string(logJson)) // push a log to task queue
	fmt.Println("queue: ", DefaultQueue.Q)

	consumer := NewConsumer()
	err := consumer.Consume()
	if err != nil {
		t.Errorf("failed: %v\n", err)
	} else {
		fmt.Println("pass")
	}
}
