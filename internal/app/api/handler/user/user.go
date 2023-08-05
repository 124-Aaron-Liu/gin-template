package user

import (
	"github.com/124-Aaron-Liu/gin-template/internal/app/api/presenter"
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

	if id == "" || fileds == "" {
		c.JSON(http.StatusOK, presenter.NewReponse("19", "INVALID_PARAMS"))
		return
	}

	// TODO: 之後重userService取
	user := []string{id, fileds}

	res := presenter.NewReponse("0", "ok").WithData(user)
	c.JSON(http.StatusOK, res)
}

func NewHandler(
	logger *zap.SugaredLogger,
) Handler {
	return Handler{logger}
}

var ProviderSet = wire.NewSet(NewHandler)
