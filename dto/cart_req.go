package dto

import "github.com/google/uuid"


type AddCartReq struct {
	UserId string `json:"user_id"`
	Items []CartReq `json:"items"`
}

type CartReq struct {
	ProductId uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

