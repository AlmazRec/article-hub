package rest

import (
	"net/http"
	"restapp/internal/messages"
	"restapp/internal/models"
	"restapp/internal/response"
	"restapp/internal/services"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthService services.AuthServiceInterface
}

func NewAuthHandler(authService services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()
	var req models.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrBadRequest,
			Error:   err.Error(),
		})
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrValidationFailed,
			Error:   err.Error(),
		})
	}

	userData, err := h.AuthService.Register(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrBadRequest,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Data:    userData,
		Message: messages.MsgRegistrationSuccess,
	})
}

func (h *AuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var req models.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrBadRequest,
			Error:   err.Error(),
		})
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: messages.ErrValidationFailed,
			Error:   err.Error(),
		})
	}

	token, err := h.AuthService.Login(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: messages.ErrInvalidCredentials,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		Data: map[string]string{
			"token": token,
		},
	})
}
