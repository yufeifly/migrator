package redis

import "github.com/yufeifly/migrator/dal"

func Get(key string) (string, error) {
	val, err := dal.GetKV(key)
	if err != nil {
		return "", err
	}
	return val, err
}
