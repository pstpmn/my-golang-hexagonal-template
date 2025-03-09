package httpHandler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/domain"
	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/port"
)

type (
	httpHandler struct {
		useUseCase port.IUserUseCase
	}

	IHttpHandler interface {
		GetAll(ctx echo.Context) error
		GetUserById(ctx echo.Context) error
	}
)

// @Summary get users
// @Description get all users from database
// @ID get-users
// @Tags Users
// @Accept json
// @Produce json
// @Router /v1/users [get]
func (h *httpHandler) GetAll(pctx echo.Context) error {
	result, err := h.useUseCase.GetAll(pctx.Request().Context())
	if err != nil {
		return ResponseError(http.StatusBadRequest, err.Error(), pctx)
	}
	return ResponseSuccess(http.StatusOK, "successful", result, pctx)
}

// @Summary get user
// @Description get users with userId from database
// @ID get-user-with-userId
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "userId"
// @Router /v1/user/{userId} [get]
func (h *httpHandler) GetUserById(pctx echo.Context) error {
	userId := pctx.Param("userId")
	result, err := h.useUseCase.GetUser(pctx.Request().Context(), userId)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return ResponseError(http.StatusNotFound, err.Error(), pctx)
		}
		return ResponseError(http.StatusBadRequest, err.Error(), pctx)
	}
	return ResponseSuccess(http.StatusOK, "successful", result, pctx)
}

func NewUserHandler(userUseCase port.IUserUseCase) IHttpHandler {
	return &httpHandler{
		useUseCase: userUseCase,
	}
}
