//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/124-Aaron-Liu/gin-template/internal/app/api"
	"github.com/124-Aaron-Liu/gin-template/internal/app/api/handler"
	"github.com/124-Aaron-Liu/gin-template/internal/app/api/handler/user"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	ProviderSet,
	api.ProviderSet,
	handler.ProviderSet,
	user.ProviderSet,
)

func NewApplication() (Application, error) {
	panic(wire.Build(providerSet))
}
