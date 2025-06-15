package main

import (
	"log"
	"my-echo-chat_service/dto"
	"my-echo-chat_service/internal/api"
	"my-echo-chat_service/internal/config"
	"my-echo-chat_service/internal/connection"
	"my-echo-chat_service/internal/repository"
	"my-echo-chat_service/internal/service"
	"net/http"

	jwtMid "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main(){
	cnf := config.LoadConfig()
	dbConnection := connection.GetDatabase()
	defer func ()  {
		if err := dbConnection.Disconnect(nil);err != nil {
			log.Println("Failed to disconnect")
		}else {
			log.Println("MongoDB disconnected")
		}
	}()


	app := echo.New()

	jwtMiddleware := jwtMid.WithConfig(jwtMid.Config{
		SigningKey: []byte(cnf.Jwt.Key),
		ErrorHandler: func (ctx echo.Context, err error) error  {
			log.Printf("error jwt: %v", err)
			return ctx.JSON(http.StatusUnauthorized, dto.CreateResponseError("Authentication failed"))
		},
	})

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowMethods: []string{echo.GET,echo.POST,echo.PUT, echo.DELETE, echo.OPTIONS},
		
	}))






	repository.NewUser(dbConnection.Database("chat_db"))
	userRepository := repository.NewUser(dbConnection.Database("chat_db"))
	userService := service.NewUser(cnf, userRepository)



	
	
	api.NewAuth(app, userService, jwtMiddleware)
	

	log.Println("🚀 Starting server on " + cnf.Server.Host + ":" + cnf.Server.Port)
	app.Logger.Fatal(app.Start(cnf.Server.Host + ":" + cnf.Server.Port))
	
}