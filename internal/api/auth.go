package api

import (
	"context"
	"log"
	"my-echo-chat_service/domain"
	"my-echo-chat_service/dto"
	"my-echo-chat_service/internal/middleware"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type authApi struct {
	authService domain.AuthService
}

func NewAuth(app *echo.Echo, authService domain.AuthService, auzMidd echo.MiddlewareFunc) {
	a := authApi{
		authService: authService,
	}

	api := app.Group("/api/auth")

	api.POST("/register", a.Register)
	api.POST("/login", a.Login)
	api.POST("/logout", a.Logout, auzMidd, middleware.CheckBlacklist)
	api.GET("/check", a.GetProfile, auzMidd, middleware.CheckBlacklist)
}

func (aa authApi) Register(ctx echo.Context) error {
	c, cancel := context.WithTimeout(ctx.Request().Context(), 10*time.Second)
	defer cancel()

	var req dto.UserData
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.CreateResponseErrorData(err.Error()))
	}

	token, err := aa.authService.Register(c, req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.CreateResponseErrorData(err.Error()))
	}

	return ctx.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Register Account", map[string]interface{}{
		"token": token,
	}))
}

func (aa authApi) Login(ctx echo.Context) error {
	c, cancel := context.WithTimeout(ctx.Request().Context(), 10*time.Second)
	defer cancel()

	var req dto.UserRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.CreateResponseErrorData(err.Error()))
	}

	res, err := aa.authService.Login(c, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return ctx.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Login Account", res))

}

func (aa authApi) GetProfile(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	userData, err := aa.authService.GetProfile(ctx.Request().Context(), userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return ctx.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Get Profile", userData))

}

func (aa authApi) Logout(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)
	tokenRaw := user.Raw

	log.Print(userId)

	if middleware.IsBlacklisted(tokenRaw) {
		return ctx.JSON(http.StatusUnauthorized, dto.CreateResponseError("You are not logged in"))
	}

	// 1. Cek logika logout di service
	err := aa.authService.Logout(ctx.Request().Context(), userId)
	if err != nil {
		log.Printf("error logout: %v", err)
		return ctx.JSON(http.StatusUnauthorized, dto.CreateResponseErrorData(err.Error()))
	}

	// 2. Tambah token ke blacklist
	middleware.AddToBlacklist(tokenRaw)

	return ctx.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Logout Account", nil))
}
