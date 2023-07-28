//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	newSugaredLogger,
	newApplication,
	newGinEngine,
	newLogger,
)

func NewApplication() (Application, error) {
	panic(wire.Build(providerSet))
}
