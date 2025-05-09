package middlewares

import (
	"net/http"
	"restapp/internal/services"
	"strings"

	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	s services.AuthServiceInterface // Используем интерфейс вместо конкретного типа
}

func NewAuthMiddleware(s services.AuthServiceInterface) *AuthMiddleware {
	return &AuthMiddleware{s: s}
}

func (h *AuthMiddleware) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Отсутствует токен авторизации"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Неверный формат токена"})
		}

		claims, err := h.s.ValidateToken(parts[1])
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}

		c.Set("user_id", claims.UserId)

		return next(c)
	}
}
