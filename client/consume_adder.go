package client

import "github.com/levigross/grequests"

// ConsumeAdder tell the proxy that I(dst) has consumed a log
func (cli *Client) ConsumeAdder() error {
	ro := &grequests.RequestOptions{}
	url := "http://127.0.0.1:6788/log/consume"
	_, err := grequests.Post(url, ro)
	if err != nil {
		return err
	}
	return nil
}
