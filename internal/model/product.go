package model

import "github.com/google/uuid"

type Product struct {
	Id    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name  string    `json:"name" gorm:"column:name"`
	Sku   string    `json:"sku" gorm:"column:sku"`
	Price float64   `json:"price" gorm:"column:price"`
	Stock int       `json:"stock" gorm:"column:stock"`
}
