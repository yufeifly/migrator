package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/scheduler"
)

//
func Set(service string, key, val string) error {
	ser, err := scheduler.DefaultScheduler.GetService(service)
	if err != nil {
		return err
	}
	err = doSetKV(ser.ServiceCli, key, val)
	if err != nil {
		return err
	}
	return nil
}

//
func doSetKV(cli *redis.Client, key, val string) error {
	err := cli.Set(context.Background(), key, val, 0).Err()
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"key":   key,
		"value": val,
	}).Info("pair set")
	return nil
}
