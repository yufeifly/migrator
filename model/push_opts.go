package model

type PushOpts struct {
	CheckpointOpts
	DestIP      string
	DestPort    string
	ContainerID string
}
