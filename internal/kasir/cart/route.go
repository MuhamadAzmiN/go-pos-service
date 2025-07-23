package cart

import (
	"my-golang-service-pos/internal/middleware"

	"github.com/labstack/echo/v4"
)


func RegisterRoutes(g *echo.Group, svc *Service, auth echo.MiddlewareFunc) {
	h := NewHandler(svc)

	r := g.Group("/cart", auth, middleware.CheckBlacklist)
	r.POST("", h.AddOrUpdate)
	r.GET("", h.GetAll)
	r.GET("/:id", h.GetCartByUserId)
	r.DELETE("/:id", h.DeleteCartById)
}