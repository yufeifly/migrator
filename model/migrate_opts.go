package model

type MigrateOpts struct {
	Address
	Container     string
	CheckpointID  string
	CheckpointDir string
}
