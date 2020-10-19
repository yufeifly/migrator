package model

type MigrateOpts struct {
	Address
	//Container     string
	ServiceID     string
	ProxyService  string
	CheckpointID  string
	CheckpointDir string
}
