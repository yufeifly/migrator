package client

import (
	"github.com/yufeifly/migrator/api/types"
	"net/http"
)

type APIClient interface {
	SendContainerCreate(options types.CreateReqOpts) ([]byte, error)
	ConsumedAdder(proxyService string) error
}

type client struct {
	addr       types.Address
	httpClient *http.Client
}

func NewClient(address types.Address) APIClient {
	return &client{
		addr:       address,
		httpClient: &http.Client{},
	}
}
