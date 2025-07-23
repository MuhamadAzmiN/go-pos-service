package main

import (
	"log"
	"my-golang-service-pos/internal/api"
	"my-golang-service-pos/internal/config"
	"my-golang-service-pos/internal/connection"
	"my-golang-service-pos/internal/kasir"
	"my-golang-service-pos/internal/repository"
	"my-golang-service-pos/internal/service"

	"github.com/labstack/echo/v4"

	myMiddleware "my-golang-service-pos/internal/middleware"

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

	jwtMiddleware := myMiddleware.NewJWTMiddleware(cnf.Jwt.Key)


	apiPath := app.Group("/api")

	apiPath.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))



	userRepository := repository.NewUser(dbGorm, db) 
	userService := service.NewUser(cnf, userRepository)


	kasir.RegisterServices(app, config.ServiceParam{
		Config: cnf,
		Db: db,
		DbGorm: dbGorm,
		JwtMiddleware: jwtMiddleware,
	})

	productRepository := repository.NewProduct(dbGorm, db)
	productService := service.NewProduct(cnf, productRepository)

	cartRepository := repository.NewCart(dbGorm, db)
	cartService := service.NewCart(cnf, cartRepository, productRepository)
	api.NewAuth(apiPath, userService, jwtMiddleware)
	api.NewProduct(apiPath, productService, jwtMiddleware)
	api.NewCart(apiPath, cartService, jwtMiddleware)

	log.Println("🚀 Starting server on " + cnf.Server.Host + ":" + cnf.Server.Port)
	app.Logger.Fatal(app.Start(cnf.Server.Host + ":" + cnf.Server.Port))

}







