package middlewares

import (
	"day-2/models"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthMiddleware(e *echo.Echo) {
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
}
