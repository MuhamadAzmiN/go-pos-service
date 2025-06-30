package middleware

import (
	"my-golang-service-pos/dto"

	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CheckBlacklist(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken := c.Get("user")
		if userToken == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token missing or invalid")
		}

		token, ok := userToken.(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token format")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid claims")
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "user_id not found in token")
		}

		c.Set("user_id", userID)

		if IsBlacklisted(token.Raw) {
			return c.JSON(http.StatusUnauthorized, dto.CreateResponseError("Token has been revoked"))
		}

		return next(c)
	}
}
