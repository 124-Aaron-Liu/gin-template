package handler

import (
	"github.com/124-Aaron-Liu/gin-template/internal/app/api/handler/user"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type Handler struct {
	logger      *zap.SugaredLogger
	UserHandler user.Handler
}

func NewHandler(
	logger *zap.SugaredLogger,
	userHandler user.Handler,
) Handler {
	return Handler{
		logger,
		userHandler,
	}
}

var ProviderSet = wire.NewSet(NewHandler)
