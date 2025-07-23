package middleware

import (
	"my-golang-service-pos/dto"

	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CheckBlacklist(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// --- ambil token dari context ---
		userToken := c.Get("user")
		if userToken == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "token missing")
		}

		token, ok := userToken.(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token format")
		}

		// --- ambil claims ---
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid claims")
		}

		// --- ambil user_id secara aman ---
		rawID, exists := claims["user_id"]
		if !exists || rawID == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "user_id missing in token")
		}

		userID, ok := rawID.(string)
		if !ok || userID == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "user_id invalid")
		}

		// simpan ke context agar bisa dipakai handler
		c.Set("user_id", userID)

		// --- cek blacklist ---
		if IsBlacklisted(token.Raw) {
			return c.JSON(http.StatusUnauthorized, dto.CreateResponseError("Token has been revoked"))
		}

		return next(c)
	}
}

