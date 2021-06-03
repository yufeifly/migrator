package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/scheduler"
)

func Delete(containerID string, key string) error {
	containerServ, err := scheduler.Default().GetContainerServ(containerID)
	if err != nil {
		return err
	}

	err = deleteKV(containerServ.ServiceCli, key)
	if err != nil {
		logrus.Errorf("delete kv failed, key: %v", key)
		return err
	}
	return nil
}

func deleteKV(cli *redis.Client, key string) error {
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
