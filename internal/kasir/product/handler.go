package product

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"my-golang-service-pos/dto"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler { return &Handler{svc: svc} }

func (h *Handler) GetList(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	data, err := h.svc.GetList(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}
	return c.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success", data))
}


func (h *Handler) GetByID(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
	data, err := h.svc.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}
	return c.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success", data))
}


func (h *Handler) Create(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	var req dto.ProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.CreateResponseErrorData(err.Error()))
	}
	if err := h.svc.Create(ctx, req); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}
	return c.JSON(http.StatusCreated, dto.CreateResponseSuccessData("Created", nil))
}

func (h *Handler) Delete(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	if err := h.svc.Delete(ctx, c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}
	return c.JSON(http.StatusOK, dto.CreateResponseSuccessData("Deleted", nil))
}
