package scheduler

import (
	"github.com/go-redis/redis/v8"
	"github.com/yufeifly/migrator/api/types/svc"
)

type Service struct {
	ID             string        // service id
	ProxyServiceID string        // proxy serviceID
	ServicePort    string        // service port, also the exposed port of the container
	ContainerID    string        // the worker container
	ServiceCli     *redis.Client // redis connection
}

// NewService new a storage service, keep it in map
func NewService(opts svc.ServiceOpts) *Service {
	return &Service{
		ID:             opts.ID,
		ProxyServiceID: opts.ProxyServiceID,
		ServicePort:    opts.ServicePort,
		ContainerID:    opts.Container,
		ServiceCli: redis.NewClient(&redis.Options{
			Addr:     "localhost" + ":" + opts.ServicePort,
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

// PseudoRegister register services
func PseudoRegister() {
	opts1 := svc.ServiceOpts{
		ID:             "service1.1",
		ProxyServiceID: "service1",
		ServicePort:    "39955",
		Container:      "redis-service1", // name of the container
	}
	DefaultScheduler.AddService(NewService(opts1))

	opts2 := svc.ServiceOpts{
		ID:             "service2.1",
		ProxyServiceID: "service2",
		ServicePort:    "39956",
		Container:      "redis-service2", // name of the container
	}
	DefaultScheduler.AddService(NewService(opts2))
}
