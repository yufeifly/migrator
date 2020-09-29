package model

type MigrateOpts struct {
	Container     string
	CheckpointID  string
	CheckpointDir string
	DestIP        string
	DestPort      string
}
