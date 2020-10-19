package scheduler

import (
	"github.com/go-redis/redis/v8"
	"github.com/yufeifly/migrator/model"
)

type Service struct {
	ID             string        // service id
	ProxyServiceID string        // proxy serviceID
	ServicePort    string        // service port, also the exposed port of the container
	ContainerID    string        // the worker container
	ServiceCli     *redis.Client // redis connection
}

func init() {
	PseudoRegister()
}

// NewService new a storage service, keep it in map
func NewService(opts model.ServiceOpts) *Service {
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
	opts1 := model.ServiceOpts{
		ID:             "service.A1",
		ProxyServiceID: "service1",
		ServicePort:    "6380",
		Container:      "9fb8d484526c",
	}
	DefaultScheduler.AddService(NewService(opts1))

	opts2 := model.ServiceOpts{
		ID:             "service.B1",
		ProxyServiceID: "service2",
		ServicePort:    "6666",
		Container:      "f1b70d823371",
	}
	DefaultScheduler.AddService(NewService(opts2))
}
