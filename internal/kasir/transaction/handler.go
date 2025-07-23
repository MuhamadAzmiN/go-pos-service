package transaction

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




func (h *Handler) GetList(c echo.Context) error {
	ctx , cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	data , err := h.svc.GetList(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return c.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Get All Transaction", data))
}

func (h *Handler) CreateTransaction(c echo.Context) error {
	ctx , cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	var req dto.TransactionRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.CreateResponseErrorData(err.Error()))
	}

	res , err := h.svc.CreateTransaction(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return c.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Create Transaction", res))


}

