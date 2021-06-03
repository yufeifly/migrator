package client

import (
	"net/http"

	"github.com/yufeifly/migrator/api/types"
	"github.com/yufeifly/migrator/api/types/log"
)

type APIClient interface {
	SendContainerCreate(options types.CreateReqOpts) ([]byte, error)
	ConsumedAdder(proxyService string) error
	SendLog(logWithID log.LogWithCID) error
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
