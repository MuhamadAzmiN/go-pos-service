package dto

import "github.com/google/uuid"

type CartItemRes struct {
	Id          uuid.UUID `json:"id"`
	ProductId   uuid.UUID `json:"product_id"`
	ProductName string    `json:"product_name"`
	Price       int   `json:"price"`
	Quantity    int       `json:"quantity"`
	Subtotal    int   `json:"subtotal"`
	Tax         string    `json:"tax"`
	Discount    string    `json:"discount"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

type CartFullRes struct {
	UserId string         `json:"user_id"`
	Items  []CartItemRes  `json:"items"`
	Total  int        `json:"total"`
}






