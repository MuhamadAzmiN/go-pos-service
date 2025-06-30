package domain

import (
	"context"
	"my-golang-service-pos/dto"
	"time"
)

type Transaction struct {
	Id             string           `json:"id"`
	InvoiceNumber  string           `json:"invoice_number"`
	UserId         string           `json:"user_id"`
	TotalPrice     float64          `json:"total_price"`
	PaidAmount     float64          `json:"paid_amount"`
	ChangeAmount   float64          `json:"change_amount"`
	PaymentMethod  string           `json:"payment_method"`
	CreatedAt      time.Time           `json:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"`
	Items          []TransactionItem `json:"items"` // relasi ke item
}

type TransactionItem struct {
	Id            string  `json:"id"`
	TransactionId string  `json:"transaction_id"`
	ProductId     string  `json:"product_id"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	Subtotal      float64 `json:"subtotal"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}






type TransactionRepository interface {
	Create(ctx context.Context, transaction Transaction, items []TransactionItem) (Transaction, error)
}

type TransactionService interface {
	CreateTransaction(ctx context.Context, req dto.TransactionRequest) (Transaction, error)
}
