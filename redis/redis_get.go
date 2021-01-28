package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/cusErr"
	"github.com/yufeifly/migrator/scheduler"
)

func Get(service string, key string) (string, error) {
	ser, err := scheduler.DefaultScheduler.GetService(service)
	if err != nil {
		return "", err
	}
	val, err := doGetKV(ser.ServiceCli, key)
	if err != nil {
		return "", err
	}
	return val, err
}

func doGetKV(cli *redis.Client, key string) (string, error) {
	header := "redis.doGetKV"
	val, err := cli.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", cusErr.ErrNotFound
	}

	if err != nil {
		logrus.Errorf("%s, err: %v", header, err)
		return "", err
	}

	logrus.WithFields(logrus.Fields{
		"key":   key,
		"value": val,
	}).Debug("the (key, value) pair")

	return val, nil
}
