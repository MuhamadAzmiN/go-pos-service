package dto

type ProductRequest struct {
	Name  string  `json:"name" validate:"required"`
	Sku   string  `json:"sku" validate:"required"`
	Price float64 `json:"price" validate:"required"`
	Stock int     `json:"stock" validate:"required"`
}


