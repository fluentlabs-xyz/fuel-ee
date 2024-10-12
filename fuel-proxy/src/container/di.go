package container

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/config"
	"github.com/olebedev/emitter"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func CreateContainer() *dig.Container {
	container := dig.New()
	must(container.Provide(config.NewConfig))
	must(container.Provide(func() *emitter.Emitter {
		return emitter.New(10)
	}))
	return container
}

func MustInvoke(container *dig.Container, function interface{}, opts ...dig.InvokeOption) {
	must(container.Invoke(function, opts...))
}

func must(err error) {
	if err != nil {
		log.Fatalf("failed to initialize DI: %s", err)
	}
}
