package client

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types"
)

// SendMigrate send migrate request to dst node
func (cli *client) SendMigrate(options types.MigrateOpts) error {
	optsJSON, err := json.Marshal(options)
	if err != nil {
		return err
	}

	ro := &grequests.RequestOptions{
		JSON: optsJSON,
	}

	//example url := "http://127.0.0.1:6789/container/migrate"
	url := cli.getAPIPath("/container/migrate")
	resp, err := grequests.Post(url, ro)
	logrus.Debug("client.SendMigrate finished")
	if err != nil {
		return err
	}
	resp.RawResponse.Body.Close()
	return nil
}
