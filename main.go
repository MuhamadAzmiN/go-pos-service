package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"my-golang-service-pos/dto"
	"my-golang-service-pos/internal/api"
	"my-golang-service-pos/internal/config"
	"my-golang-service-pos/internal/connection"
	"my-golang-service-pos/internal/repository"
	"my-golang-service-pos/internal/service"
	"net/http"

	jwtMid "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cnf := config.LoadConfig()
	db, dbGorm, err := connection.NewDb(cnf, "postgres")
	if err != nil {
		log.Fatal("❌ Failed to connect to the database:", err)
	}

	if db == nil {
		log.Fatal("❌ Failed to connect to the database")
	}

	app := echo.New()

	jwtMiddleware := jwtMid.WithConfig(jwtMid.Config{
		SigningKey: []byte(cnf.Jwt.Key),
		ContextKey: "user",
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



	userRepository := repository.NewUser(dbGorm, db) 
	userService := service.NewUser(cnf, userRepository)

	productRepository := repository.NewProduct(dbGorm, db)
	productService := service.NewProduct(cnf, productRepository)

	cartRepository := repository.NewCart(dbGorm, db)
	cartService := service.NewCart(cnf, cartRepository, productRepository)

	transactionRepository := repository.NewTransaction(dbGorm, db)
	transactionService := service.NewTransaction(cnf, transactionRepository, cartRepository, productRepository)

	
	api.NewTransaction(apiPath, transactionService, jwtMiddleware)
	api.NewAuth(apiPath, userService, jwtMiddleware)
	api.NewProduct(apiPath, productService, jwtMiddleware)
	api.NewCart(apiPath, cartService, jwtMiddleware)

	log.Println("🚀 Starting server on " + cnf.Server.Host + ":" + cnf.Server.Port)
	app.Logger.Fatal(app.Start(cnf.Server.Host + ":" + cnf.Server.Port))

}




