package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID             uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	InvoiceNumber  string     `json:"invoice_number"`       // Contoh: TRX-20250630-0001
	UserID         uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	User           User       `gorm:"foreignKey:UserID" json:"-"` // Optional: relasi ke kasir

	TotalPrice     float64    `json:"total_price"`          // Total belanja
	PaidAmount     float64    `json:"paid_amount"`          // Uang yang dibayarkan
	ChangeAmount   float64    `json:"change_amount"`        // Kembalian
	PaymentMethod  string     `json:"payment_method"`       // Contoh: cash, qris, debit

	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`

	Items          []TransactionItems `gorm:"foreignKey:TransactionID" json:"items"` // relasi ke item transaksi
}
