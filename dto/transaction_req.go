package dto

import "github.com/google/uuid"

type TransactionRequest struct {
	UserId uuid.UUID `json:"user_id"`
	PaidAmount float64 `json:"paid_amount"`
	PaymentMethod string `json:"payment_method"`
}

