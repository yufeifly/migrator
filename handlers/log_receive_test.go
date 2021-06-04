package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types"
	"github.com/yufeifly/migrator/api/types/log"
	"github.com/yufeifly/migrator/client"
	"testing"
)

func TestReceiveLog(t *testing.T) {
	logWithCID := log.LogWithCID{
		CID: "s1.c1",
		Log: log.Log{
			Last:     true,
			LogQueue: []log.KV{log.KV{Key: "k1", Val: "v1"}},
		},
	}

	cli := client.NewClient(types.Address{
		IP:   "127.0.0.1",
		Port: "6789",
	})

	err := cli.SendLog(logWithCID)
	if err != nil {
		logrus.Errorf("client.SendLog Post failed, err: %v", err)
		return
	}
}
