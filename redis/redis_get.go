package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/cuserr"
	"github.com/yufeifly/migrator/scheduler"
)

func Get(service string, key string) (string, error) {
	containerServ, err := scheduler.Default().GetContainerServ(service)
	if err != nil {
		return "", err
	}
	val, err := getKV(containerServ.ServiceCli, key)
	if err != nil {
		return "", err
	}
	return val, err
}

func getKV(cli *redis.Client, key string) (string, error) {
	header := "redis.getKV"
	val, err := cli.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", cuserr.ErrNotFound
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
