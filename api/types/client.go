package types

type CreateOpts struct {
	ContainerName string
	ImageName     string
	HostPort      string
	ContainerPort string
	PortBindings  string
	ExposedPorts  string
	Cmd           string
}

type CreateReqOpts struct {
	CreateOpts
	Address
}
