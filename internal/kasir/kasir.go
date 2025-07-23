package kasir

import (
	"my-golang-service-pos/internal/config"
	"my-golang-service-pos/internal/kasir/cart"
	"my-golang-service-pos/internal/kasir/product"
	"my-golang-service-pos/internal/kasir/transaction"
	"my-golang-service-pos/internal/repository"

	"github.com/labstack/echo/v4"
)

func RegisterServices(e *echo.Echo, p config.ServiceParam) {
	productRepo := repository.NewProduct(p.DbGorm, p.Db)
	productSvc := product.NewService(p.Config, productRepo)


	cartRepo := repository.NewCart(p.DbGorm, p.Db)
	cartSvc := cart.NewService(p.Config, cartRepo, productRepo)



	transactionRepo := repository.NewTransaction(p.DbGorm, p.Db)
	transactionSvc := transaction.NewService(p.Config, transactionRepo, cartRepo, productRepo)

	group := e.Group("/v1/api/kasir")
	product.RegisterRoutes(group, productSvc, p.JwtMiddleware)
	transaction.RegisterRoutes(group, transactionSvc, p.JwtMiddleware)
	cart.RegisterRoutes(group, cartSvc, p.JwtMiddleware)

}

