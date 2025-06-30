package repository

import (
	"context"
	"my-golang-service-pos/domain"
	"my-golang-service-pos/internal/iface"
	"time"

	"github.com/google/uuid"
)



type transactionRepository struct {
	dbGorm iface.IGorm
	db iface.ISqlx
}


func NewTransaction(dbGorm iface.IGorm, db iface.ISqlx) domain.TransactionRepository {
	return &transactionRepository{
		dbGorm: dbGorm,
		db: db,
	}
}



func (t transactionRepository) Create(ctx context.Context, transaction domain.Transaction, items []domain.TransactionItem) (domain.Transaction, error) {
	tx := t.dbGorm.Begin()

	// 1. Simpan transaksi
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return domain.Transaction{}, err
	}

	// 2. Loop item dan pastikan valid
	for i := range items {
		if items[i].Id == "" {
			items[i].Id = uuid.New().String()
		}
		if items[i].CreatedAt == "" {
			items[i].CreatedAt = time.Now().Format(time.RFC3339)
		}
		if items[i].UpdatedAt == "" {
			items[i].UpdatedAt = time.Now().Format(time.RFC3339)
		}
		items[i].TransactionId = transaction.Id
	}

	// 3. Simpan semua item
	if err := tx.Create(&items).Error; err != nil {
		tx.Rollback()
		return domain.Transaction{}, err
	}

	// 4. Commit transaksi
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return domain.Transaction{}, err
	}

	return transaction, nil
}
