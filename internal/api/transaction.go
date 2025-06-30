package api

import (
	"context"
	"my-golang-service-pos/domain"
	"my-golang-service-pos/dto"
	"my-golang-service-pos/internal/middleware"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type transactionApi struct {
	transactionService domain.TransactionService
}


func NewTransaction(app *echo.Group, transactionService domain.TransactionService, auzMidd echo.MiddlewareFunc) {
	a := transactionApi{
		transactionService: transactionService,
	}

	api := app.Group("/transaction")

	api.POST("/", a.CreateTransaction, auzMidd, middleware.CheckBlacklist)

}


func (tr transactionApi) CreateTransaction(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	var req dto.TransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.CreateResponseErrorData(err.Error()))
	}	

	res, err := tr.transactionService.CreateTransaction(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateResponseErrorData(err.Error()))
	}

	return c.JSON(http.StatusOK, dto.CreateResponseSuccessData("Success Create Transaction", res))
}