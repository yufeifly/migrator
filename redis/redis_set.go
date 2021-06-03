package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/cuserr"
	"github.com/yufeifly/migrator/scheduler"
)

//
func Set(containerID string, key, val string) error {
	cServ, err := scheduler.Default().GetContainerServ(containerID)
	if err != nil {
		logrus.Errorf("redis.GetService err : %v", err)
		return err
	}

	token := cServ.Ticket().Get()
	if cServ.Ticket().WriteBaned(token) {
		return cuserr.ErrServiceNotAvailable
	}
	if cServ.Ticket().IsLogging(token) {
		go func(key, val string) {
			err := cServ.LogRecord(key, val)
			if err != nil {
				logrus.Errorf("redis.Set LogRecord failed, err: %v", err)
			}
		}(key, val)
	}

	err = setKV(cServ.ServiceCli, key, val)
	if err != nil {
		return err
	}
	return nil
}

//
func setKV(cli *redis.Client, key, val string) error {
	err := cli.Set(context.Background(), key, val, 0).Err()
	if err != nil {
		logrus.Errorf("redis.doSetKV err : %v", err)
		return err
	}
	logrus.WithFields(logrus.Fields{
		"key":   key,
		"value": val,
	}).Debug("pair set")
	return nil
}
