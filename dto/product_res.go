package dto

type ProductResponse struct {
	Id    string  `json:"id" validate:"required"`
	Name  string  `json:"name"`
	Sku   string  `json:"sku"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
