package dto

type ProductData struct {
	Id    string  `json:"id" validate:"required"`
	Name  string  `json:"name" validate:"required"`
	Sku   string  `json:"sku" validate:"required"`
	Price int `json:"price" validate:"required"`
	Stock int     `json:"stock" validate:"required"`
}


