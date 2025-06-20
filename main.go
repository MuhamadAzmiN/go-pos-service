package main

import (
	"log"
	"my-golang-service-pos/dto"
	"my-golang-service-pos/internal/api"
	"my-golang-service-pos/internal/config"
	"my-golang-service-pos/internal/connection"
	"my-golang-service-pos/internal/repository"
	"my-golang-service-pos/internal/service"
	"net/http"
	"github.com/labstack/echo/v4"

	jwtMid "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cnf := config.LoadConfig()
	db := connection.InitPostgres(cnf)
	
	if db == nil {
		log.Fatal("❌ Failed to connect to the database")
	}

	app := echo.New()

	jwtMiddleware := jwtMid.WithConfig(jwtMid.Config{
		SigningKey: []byte(cnf.Jwt.Key),
		ErrorHandler: func(ctx echo.Context, err error) error {
			return ctx.JSON(http.StatusUnauthorized, dto.CreateResponseError("Authentication failed"))
		},
	})

	apiPath := app.Group("/api")

	apiPath.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))

	userRepository := repository.NewUser(db)
	userService := service.NewUser(cnf, userRepository)

	api.NewAuth(apiPath, userService, jwtMiddleware)

	log.Println("🚀 Starting server on " + cnf.Server.Host + ":" + cnf.Server.Port)
	app.Logger.Fatal(app.Start(cnf.Server.Host + ":" + cnf.Server.Port))

}

