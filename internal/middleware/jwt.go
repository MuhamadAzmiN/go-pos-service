package middleware

import (
	"my-golang-service-pos/dto"
	"net/http"

	"github.com/labstack/echo/v4"
	jwtMid "github.com/labstack/echo-jwt/v4"
)

// NewJWTMiddleware mengembalikan middleware JWT siap pakai
func NewJWTMiddleware(secret string) echo.MiddlewareFunc {
	return jwtMid.WithConfig(jwtMid.Config{
		SigningKey: []byte(secret),
		ContextKey: "user",
		ErrorHandler: func(ctx echo.Context, err error) error {
			return ctx.JSON(http.StatusUnauthorized, dto.CreateResponseError("Authentication failed"))
		},
	})
}
