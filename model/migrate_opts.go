package model

type MigrateOpts struct {
	ContainerName string
	CheckpointID  string
	CheckpointDir string
	DestIP        string
	DestPort      string
}
