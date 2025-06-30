package api

import (

	"my-golang-service-pos/domain"
	"my-golang-service-pos/dto"
	"my-golang-service-pos/internal/middleware"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

type cartApi struct {
	cartService domain.CartService
}


func NewCart(app *echo.Group, cartService domain.CartService, auzMidd echo.MiddlewareFunc) {
	a := cartApi{
		cartService: cartService,
	}

	api := app.Group("/cart")

	api.POST("/", a.AddOrUpdate, auzMidd, middleware.CheckBlacklist)
	api.GET("/list", a.GetAll, auzMidd, middleware.CheckBlacklist)
	api.GET("/", a.GetCartByUserId, auzMidd, middleware.CheckBlacklist)
	api.DELETE("/:id", a.DeleteCartById, auzMidd, middleware.CheckBlacklist)
}






func (cc cartApi) AddOrUpdate(ctx echo.Context) error {
	c, cancel := context.WithTimeout(ctx.Request().Context(), 10*time.Second)
	defer cancel()

	var req dto.AddCartReq
	userIDRaw := ctx.Get("user_id").(string)
	req.UserId = userIDRaw

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.CreateResponseErrorData(err.Error()))
	}

	err := cc.cartService.AddOrUpdate(c, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return ctx.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Add Or Update Cart", nil))

}



func (cc cartApi) GetAll(ctx echo.Context) error {
	c, cancel := context.WithTimeout(ctx.Request().Context(), 10*time.Second)
	defer cancel()

	res, err := cc.cartService.GetAll(c)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return ctx.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Get All Cart", res))
}


func (cc cartApi) GetCartByUserId(ctx echo.Context) error {
	c, cancel := context.WithTimeout(ctx.Request().Context(), 10*time.Second)
	defer cancel()

	userID := ctx.Get("user_id").(string)
	
	res, err := cc.cartService.GetCartByUserId(c, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return ctx.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Get Cart By Id", res))
}



func (cc cartApi) DeleteCartById(c echo.Context) error {
	ctx , cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
	err := cc.cartService.DeleteCartById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return c.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Delete Cart By Id", nil))
}


