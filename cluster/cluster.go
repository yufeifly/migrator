package cluster

import "github.com/yufeifly/migrator/api/types"

var defaultCluster cluster

type Cluster interface {
	GetProxy() Node
}

type cluster struct {
	Master Node `json:"proxy"`
}

type Node struct {
	types.Address
}

func DefaultCluster() Cluster {
	return &defaultCluster
}

func (c *cluster) GetProxy() Node {
	return c.Master
}
