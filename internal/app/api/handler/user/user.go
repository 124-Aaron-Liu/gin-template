package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	logger *zap.SugaredLogger
}

func (h Handler) GetUser(c *gin.Context) {

	id := c.Query("id")
	fileds := c.Query("field")

	c.JSON(http.StatusOK, []string{id, fileds})
}

func NewHandler(
	logger *zap.SugaredLogger,
) Handler {
	return Handler{logger}
}

var ProviderSet = wire.NewSet(NewHandler)
