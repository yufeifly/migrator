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
