package model

type MigrateOpts struct {
	Address
	Container     string
	ServiceID     string
	CheckpointID  string
	CheckpointDir string
}
