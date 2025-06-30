package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionItems struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`

	TransactionID uuid.UUID  `gorm:"type:uuid" json:"transaction_id"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID" json:"-"` // Optional: relasi ke transaksi

	ProductID     uuid.UUID  `gorm:"type:uuid" json:"product_id"`
	Product       Product     `gorm:"foreignKey:ProductID" json:"-"`     // Optional: relasi ke produk

	Quantity      int        `json:"quantity"`
	Price         float64    `json:"price"`      // Harga satuan saat transaksi
	Subtotal      float64    `json:"subtotal"`   // Price * Quantity

	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
