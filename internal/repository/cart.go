package repository

import (
	"context"
	"my-golang-service-pos/domain"
	"my-golang-service-pos/internal/iface"

	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)



type cartRepository struct {
	dbGorm iface.IGorm
	db iface.ISqlx
}


func NewCart(dbGorm iface.IGorm, db iface.ISqlx) domain.CartRepository {
	return &cartRepository{
		dbGorm: dbGorm,
		db: db,
	}
}

func (c cartRepository) Insert(ctx context.Context, userId string, productId uuid.UUID, quantity int) error {
	var cart domain.Cart

	err := c.dbGorm.WithContext(ctx).
		Where("product_id = ? AND user_id = ?", productId, userId).
		Take(&cart).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Cart belum ada, insert baru
			newCart := domain.Cart{
				Id:        uuid.New(),
				UserId:    userId,
				ProductId: productId,
				Quantity:  quantity,
			}
			return c.dbGorm.WithContext(ctx).Create(&newCart).Error
		}

		return err
	}

	// Cart sudah ada, update quantity
	cart.Quantity += quantity
	return c.dbGorm.WithContext(ctx).Save(&cart).Error
}


func (c cartRepository) GetByCartId(ctx context.Context, userId string) ([]domain.Cart, error) {
	var dataset []domain.Cart
	err := c.dbGorm.WithContext(ctx).
		Preload("Product").
		Where("user_id = ?", userId).
		Order("created_at ASC").
		Find(&dataset).Error

	if err != nil {
		return nil, err
	}

	return dataset, nil
}


func (c cartRepository) GetAll(ctx context.Context) ([]domain.Cart, error) {
	dataset := []domain.Cart{}
	err := c.dbGorm.WithContext(ctx).Preload("Product").Find(&dataset).Error
	if err != nil {
		return nil, err
	}

	return dataset, nil
}


func (c cartRepository) Delete(ctx context.Context, id string) error {
	err := c.dbGorm.WithContext(ctx).Where("id = ?", id).Delete(&domain.Cart{}).Error
	if err != nil {
		return err
	}

	return nil
}





func (c cartRepository) ClearCart(ctx context.Context, userId string) error {
	err := c.dbGorm.WithContext(ctx).Where("user_id = ?", userId).Delete(&domain.Cart{}).Error
	if err != nil {
		return err
	}

	return nil
}