package api

import (
	"github.com/124-Aaron-Liu/gin-template/internal/app/api/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func RouterWrap(h handler.Handler) func(r *gin.Engine) {
	return func(r *gin.Engine) {
		r.GET("/user", h.UserHandler.GetUser)
	}
}

var ProviderSet = wire.NewSet(RouterWrap)
