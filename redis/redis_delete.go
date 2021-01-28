package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/scheduler"
)

func Delete(service string, key string) error {
	ser, err := scheduler.DefaultScheduler.GetService(service)
	if err != nil {
		return err
	}
	logrus.Debugf("redis.service: %v", ser)
	err = doDeleteKV(ser.ServiceCli, key)
	if err != nil {
		return err
	}
	return nil
}

func doDeleteKV(cli *redis.Client, key string) error {
	err := cli.Del(context.Background(), key).Err()
	if err != nil {
		logrus.Errorf("redis.doDeleteKV err : %v", err)
		return err
	}
	logrus.WithFields(logrus.Fields{
		"key": key,
	}).Debug("pair delete")
	return nil
}
