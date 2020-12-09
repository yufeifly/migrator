package client

import "fmt"

// getAPIPath path means webapi path, for example: /redis/set
func (cli *client) getAPIPath(path string) string {
	return fmt.Sprintf("http://%s:%s%s", cli.addr.IP, cli.addr.Port, path)
}
