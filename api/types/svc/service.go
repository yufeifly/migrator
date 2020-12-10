package svc

type ServiceOpts struct {
	ID             string // service id
	ProxyServiceID string
	ServicePort    string
	Container      string
}
