package user

import (
	"github.com/124-Aaron-Liu/gin-template/internal/app/api/presenter"
	userSvc "github.com/124-Aaron-Liu/gin-template/internal/app/usecase/user"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handler struct {
	logger      *zap.SugaredLogger
	userService userSvc.UseCase
}

func (h Handler) GetUser(c *gin.Context) {
	idStr := c.Query("id")
	fileds := c.Query("field")

	if idStr == "" || fileds == "" {
		c.JSON(http.StatusOK, presenter.NewReponse("19", "INVALID_PARAMS"))
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, presenter.NewReponse("19", "INVALID_PARAMS"))
	}

	user, _ := h.userService.GetUserById(c, uint32(id))

	res := presenter.NewReponse("0", "ok").WithData(user)
	c.JSON(http.StatusOK, res)
}

func NewHandler(
	logger *zap.SugaredLogger,
	userService userSvc.UseCase,
) Handler {
	return Handler{logger, userService}
}

var ProviderSet = wire.NewSet(NewHandler)
