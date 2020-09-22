package container

import (
	"github.com/fgrosse/goldi"

	SharedConfig "github.com/github-profile/go-nats/shared/config"
	SharedNats "github.com/github-profile/go-nats/shared/nats"
)

// DefaultContainer returns default given depedency injections
func DefaultContainer() *goldi.Container {
	registry := goldi.NewTypeRegistry()

	config := make(map[string]interface{})
	container := goldi.NewContainer(registry, config)

	container.RegisterType("shared.config", SharedConfig.NewImmutableConfig)
	container.RegisterType("shared.nats", SharedNats.NewService, "@shared.config")
	return container
}
