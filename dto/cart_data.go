package dto

import "github.com/google/uuid"

type CartData struct {
	ProductId uuid.UUID `json:"productId"`
	Quantity  int       `json:"quantity"`
}