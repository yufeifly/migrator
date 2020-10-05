package client

import "net/http"

type Client struct {
	Client *http.Client
}

func NewClient() *Client {
	return &Client{
		Client: &http.Client{},
	}
}
