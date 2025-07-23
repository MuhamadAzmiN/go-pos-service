package cart

import (
	"context"
	"my-golang-service-pos/dto"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc *Service
}


func NewHandler(svc *Service) *Handler { return &Handler{svc: svc} }



func (h *Handler) AddOrUpdate(ctx echo.Context) error {
	c, cancel := context.WithTimeout(ctx.Request().Context(), 10*time.Second)
	defer cancel()

	var req dto.AddCartReq
	userIDRaw := ctx.Get("user_id").(string)
	req.UserId = userIDRaw

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.CreateResponseErrorData(err.Error()))
	}

	err := h.svc.AddOrUpdate(c, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return ctx.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Add Or Update Cart", nil))

}



func (h *Handler) GetAll(ctx echo.Context) error {
	c, cancel := context.WithTimeout(ctx.Request().Context(), 10*time.Second)
	defer cancel()

	res, err := h.svc.GetAll(c)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return ctx.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Get All Cart", res))
}


func (h *Handler) GetCartByUserId(ctx echo.Context) error {
	c, cancel := context.WithTimeout(ctx.Request().Context(), 10*time.Second)
	defer cancel()

	userID := ctx.Get("user_id").(string)
	
	res, err := h.svc.GetCartByUserId(c, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return ctx.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Get Cart By Id", res))
}



func (h *Handler) DeleteCartById(c echo.Context) error {
	ctx , cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
	err := h.svc.DeleteCartById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return c.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Delete Cart By Id", nil))
}


