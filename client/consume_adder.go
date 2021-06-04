package client

import (
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
)

// ConsumeAdder tell the src node that I(dst) have consumed a log
func (cli *client) ConsumedAdder(cid string) error {
	data := make(map[string]string)
	data["ContainerID"] = cid
	ro := &grequests.RequestOptions{
		Data: data,
	}

	url := cli.getAPIPath("/log/consume")
	_, err := grequests.Post(url, ro)
	if err != nil {
		logrus.Errorf("client.ConsumedAdder: post err %v", err)
		return err
	}
	return nil
}
