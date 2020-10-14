package task

import (
	"encoding/json"
	"fmt"
	"github.com/yufeifly/migrator/model"
	"strconv"
	"testing"
)

// TestConsumer_Consume two logs
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
	log.SetLastFlag(false)

	fmt.Printf("log: %v\n", log)

	logJson, _ := json.Marshal(log)
	DefaultQueue.Push(string(logJson)) // push a log to task queue

	var data2 []string
	for i := 50; i < 100; i++ {
		s := strconv.Itoa(i)
		tmp := []string{s, s + "#"}
		tmpJson, _ := json.Marshal(tmp)
		data2 = append(data2, string(tmpJson))
	}

	log.SetLogQueue(data2)
	log.SetLastFlag(true)

	fmt.Printf("log: %v\n", log)

	logJson2, _ := json.Marshal(log)
	DefaultQueue.Push(string(logJson2)) // push last log to task queue

	//
	consumer := NewConsumer()
	err := consumer.Consume()
	if err != nil {
		t.Errorf("failed: %v\n", err)
	} else {
		fmt.Println("pass")
	}
}
