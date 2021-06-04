package client

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types/log"
)

// SendLog send a log to dst,
func (cli *client) SendLog(logWithCID log.LogWithCID) error {
	logrus.Debugf("data to send: %v", logWithCID.Log)
	dataJSON, err := json.Marshal(logWithCID)
	if err != nil {
		logrus.Errorf("client.SendLog Marshal failed, err :%v", err)
		return err
	}

	ro := &grequests.RequestOptions{
		JSON: dataJSON,
	}

	//url := "http://" + cli.Target.IP + ":" + cli.Target.Port + "/logger"
	url := cli.getAPIPath("/logger")
	resp, err := grequests.Post(url, ro)
	if err != nil {
		logrus.Errorf("client.SendLog Post failed, err: %v", err)
		return err
	}
	logrus.Infof("client.SendLog resp: %v", resp)
	resp.RawResponse.Body.Close()
	return nil
}
