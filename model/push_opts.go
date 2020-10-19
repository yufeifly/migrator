package model

type PushOpts struct {
	//CheckpointOpts
	CheckPointID  string
	CheckPointDir string
	DestIP        string
	DestPort      string
	ContainerID   string
	ServiceID     string
	ServicePort   string
	ProxyService  string
}
