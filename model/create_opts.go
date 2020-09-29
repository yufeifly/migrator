package model

type CreateOpts struct {
	ContainerName string
	ImageName     string
	HostPort      string
	ContainerPort string
	PortBindings  string
	ExposedPorts  string
	Cmd           string
	//DestIP        string // for sending create request
}
