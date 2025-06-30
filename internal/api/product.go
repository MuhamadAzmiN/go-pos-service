package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"my-golang-service-pos/domain"
	"my-golang-service-pos/dto"
	"my-golang-service-pos/internal/middleware"
	"net/http"
	"time"
)

type productApi struct {
	productService domain.ProductService
}

func NewProduct(app *echo.Group, productService domain.ProductService, auzMidd echo.MiddlewareFunc) {
	a := productApi{
		productService: productService,
	}

	api := app.Group("/product")

	api.GET("/list", a.GetProductList)
	api.GET("/:id", a.GetProductById, auzMidd, middleware.CheckBlacklist)
	api.POST("/", a.CreateProduct, auzMidd, middleware.CheckBlacklist)
	api.DELETE("/:id", a.DeleteProduct, auzMidd, middleware.CheckBlacklist)

}

func (pr productApi) GetProductList(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	res, err := pr.productService.GetProductList(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return c.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Get Product List", res))
}

func (pr productApi) GetProductById(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
	res, err := pr.productService.GetProductById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return c.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Get Product By Id", res))
}

func (pr productApi) CreateProduct(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	var req dto.ProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.CreateResponseErrorData(err.Error()))
	}

	err := pr.productService.CreateProduct(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return c.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Create Product", nil))
}


func (pr productApi) DeleteProduct(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()
	

	id := c.Param("id")
	err := pr.productService.DeleteProduct(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return c.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Delete Product", nil))
}