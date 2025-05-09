package rest

import (
	"net/http"
	"restapp/internal/errors"
	"restapp/internal/messages"
	"restapp/internal/models"
	"restapp/internal/services"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthService services.AuthServiceInterface
}

func NewAuthHandler(AuthService services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{AuthService: AuthService}
}

func (h *AuthHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()
	var req models.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewErrorResponse(
			http.StatusBadRequest,
			messages.ErrBadRequest,
			err.Error(),
		))
	}

	if err := h.AuthService.Register(ctx, &req); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewErrorResponse(
			http.StatusBadRequest,
			messages.ErrBadRequest,
			err.Error(),
		))
	}

	return c.JSON(http.StatusCreated, messages.MsgRegistrationSuccess)
}

func (h *AuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var req models.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewErrorResponse(
			http.StatusBadRequest,
			messages.ErrBadRequest,
			err.Error(),
		))
	}

	token, err := h.AuthService.Login(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, errors.NewErrorResponse(
			http.StatusUnauthorized,
			messages.ErrInvalidCredentials,
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
