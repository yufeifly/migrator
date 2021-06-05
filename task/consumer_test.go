package task

import (
	"fmt"
	"github.com/yufeifly/migrator/api/types"
	"github.com/yufeifly/migrator/api/types/log"
	"github.com/yufeifly/migrator/cluster"
	"github.com/yufeifly/migrator/scheduler"
	"testing"
)

// TestConsumer_Consume two logs
func TestConsumer_Consume(t *testing.T) {
	// register a redis container service
	scheduler.Init()
	scheduler.RegisterServices()

	// send a log from src node
	logWithCID := log.LogWithCID{
		CID: "s1.c1",
		Log: log.Log{
			Last:     false,
			LogQueue: []log.KV{{Key: "k1", Val: "v1"}, {Key: "k2", Val: "v2"}},
		},
	}
	logWithCID2 := log.LogWithCID{
		CID: "s1.c1",
		Log: log.Log{
			Last:     true,
			LogQueue: []log.KV{{Key: "k3", Val: "v3"}, {Key: "k4", Val: "v4"}},
		},
	}

	newt := NewTask(logWithCID.CID)
	Default().AddTask(logWithCID.CID, newt)

	Default().GetTask(logWithCID.CID).Push(logWithCID.Log)
	Default().GetTask(logWithCID2.CID).Push(logWithCID2.Log)
	// consume the log
	consumer := NewConsumer()
	err := consumer.Consume("s1.c1", cluster.Node{Address: types.Address{IP: "127.0.0.1", Port: "6789"}})
	if err != nil {
		t.Errorf("failed: %v\n", err)
	} else {
		fmt.Println("pass")
	}
}
