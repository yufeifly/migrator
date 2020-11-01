package model

// MigrateOpts
type MigrateOpts struct {
	Address
	ServiceID     string
	ProxyService  string
	CheckpointID  string
	CheckpointDir string
}

// PushOpts push checkpoint to dst node
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
