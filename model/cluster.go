package model

type Cluster struct {
	Master Node `json:"proxy"`
}

type Node struct {
	Address
}

func (c *Cluster) GetProxy() Node {
	return c.Master
}
