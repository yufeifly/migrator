package types

// CreateOpts ...
type CreateOpts struct {
	ContainerName string
	ImageName     string
	HostPort      string
	ContainerPort string
	PortBindings  string
	ExposedPorts  string
	Cmd           string
}

// CreateReqOpts ...
type CreateReqOpts struct {
	CreateOpts
	Address
}

// CheckpointReqOpts ...
type CheckpointReqOpts struct {
	Container     string
	CheckPointID  string
	CheckPointDir string
}

// MigrateOpts
type MigrateOpts struct {
	Address       `json:"target"`
	SID           string `json:"sid"`
	CID           string `json:"c_name"`
	CheckpointID  string `json:"checkpoint_id"`
	CheckpointDir string `json:"checkpoint_dir"`
}
