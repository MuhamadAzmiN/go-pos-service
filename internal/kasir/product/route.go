package product

import (
	"my-golang-service-pos/internal/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, svc *Service, auth echo.MiddlewareFunc) {
	h := NewHandler(svc)
	r := g.Group("/products", auth, middleware.CheckBlacklist)

	r.GET("", h.GetList)
	r.GET("/:id", h.GetByID)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
}
