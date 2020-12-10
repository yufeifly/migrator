package client

import (
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
)

// ConsumeAdder tell the proxy that I(dst) has consumed a log
func (cli *client) ConsumedAdder(proxyService string) error {
	data := make(map[string]string)
	data["ProxyServiceID"] = proxyService
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
