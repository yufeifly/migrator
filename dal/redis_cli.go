package dal

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/cusErr"
)

var (
	rdb *redis.Client
)
var ctx = context.Background()

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func SetKV(key, val string) error {
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"key":   key,
		"value": val,
	}).Info("pair set")
	return nil
}

func GetKV(key string) (string, error) {
	header := "dal.GetKV"
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", cusErr.ErrNotFound
	} else if err != nil {
		logrus.Errorf("%s, err: %v", header, err)
		return "", err
	} else {
		logrus.WithFields(logrus.Fields{
			"key":   key,
			"value": val,
		}).Info("the (key, value) pair")
	}
	return val, nil
}
