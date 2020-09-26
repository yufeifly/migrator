package model

type MigrationOpts struct {
	CheckpointOpts
	DestIP   string
	DestPort string
}
