package redis

import "github.com/yufeifly/migrator/dal"

func Set(key, val string) error {
	err := dal.SetKV(key, val)
	if err != nil {
		return err
	}
	return nil
}
