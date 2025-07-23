package transaction

import (
	"my-golang-service-pos/internal/middleware"

	"github.com/labstack/echo/v4"
)


func RegisterRoutes(g *echo.Group, svc *Service, auth echo.MiddlewareFunc) {
	h := NewHandler(svc)

	r := g.Group("/transaction", auth, middleware.CheckBlacklist)
	r.POST("/", h.CreateTransaction)
	r.GET("/list", h.GetList)

}




