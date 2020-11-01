package client

import (
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
)

// ConsumeAdder tell the proxy that I(dst) has consumed a log
func (cli *Client) ConsumeAdder() error {
	header := "client.ConsumeAdder"
	ro := &grequests.RequestOptions{}
	url := "http://127.0.0.1:6788/log/consume"
	_, err := grequests.Post(url, ro)
	if err != nil {
		logrus.Errorf("%s: post err %v", header, err)
		return err
	}
	return nil
}
