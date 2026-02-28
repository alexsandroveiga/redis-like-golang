//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"

	"github.com/alexsandroveiga/redis-like-golang/internal/adapter/handler"
	"github.com/alexsandroveiga/redis-like-golang/internal/adapter/protocol"
	"github.com/alexsandroveiga/redis-like-golang/internal/infra/persistence"
	"github.com/alexsandroveiga/redis-like-golang/internal/infra/storage"
	"github.com/alexsandroveiga/redis-like-golang/internal/usecase"
)

func InitializeContainer(opt persistence.AOFProviderOption) (*Container, func(), error) {
	wire.Build(
		// Infraestructure providers
		storage.NewStore,
		persistence.NewAOFProvider,

		// Adapter providers
		protocol.NewParser,

		// Use case providers
		usecase.NewStats,
		usecase.NewCommandHandler,

		// Handler providers
		handler.NewTCPHandler,

		// Container provider
		NewContainer,
	)
	return nil, nil, nil
}
