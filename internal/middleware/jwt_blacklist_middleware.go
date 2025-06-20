package middleware

import (
	"my-golang-service-pos/dto"

	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTBlacklistMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		if user == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		token := user.(*jwt.Token).Raw

		if IsBlacklisted(token) {
			return c.JSON(http.StatusUnauthorized, dto.CreateResponseError("Token has been revoked"))
		}

		return next(c)
	}
}

func CheckBlacklist(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		tokenString := user.Raw

		if IsBlacklisted(tokenString) {
			return c.JSON(http.StatusUnauthorized, dto.CreateResponseError("Token has been revoked"))
		}

		return next(c)
	}
}
