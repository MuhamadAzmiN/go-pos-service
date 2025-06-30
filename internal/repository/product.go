package repository

import (
	"context"
	"my-golang-service-pos/domain"
	"my-golang-service-pos/dto"
	"my-golang-service-pos/internal/iface"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepository struct {
	dbGorm iface.IGorm
	db     iface.ISqlx
}


func NewProduct(dbGorm iface.IGorm, db iface.ISqlx) domain.ProductRepository {
	return &productRepository{
		dbGorm: dbGorm,
		db:     db,
	}    
}




func (p productRepository) FindAll(ctx context.Context) (result []dto.ProductResponse, err error) {
	dataset := []dto.ProductResponse{}
	err = p.dbGorm.Model(&domain.Product{}).Find(&dataset).Error
	if err != nil {
		return nil, err
	}

	return dataset, nil
}

func (p productRepository) FindById(ctx context.Context, id string) (domain.Product, error) {
	dataset := domain.Product{}
	err := p.dbGorm.Model(&domain.Product{}).Where("id = ?", id).First(&dataset).Error
	if err != nil {
		return domain.Product{}, err
	}

	return dataset, nil
}



func (p productRepository) Insert(ctx context.Context, req dto.ProductRequest) error {
	product := domain.Product{
		Id:    uuid.New(),
		Name:  req.Name,
		Sku:   req.Sku,
		Price: int(req.Price),
		Stock: req.Stock,
	}

	err := p.dbGorm.WithContext(ctx).Create(&product).Error
	if err != nil {
		return err
	}

	return nil
}


func (p productRepository) Delete(ctx context.Context, id string) error {
	err := p.dbGorm.WithContext(ctx).Where("id = ?", id).Delete(&domain.Product{}).Error
	if err != nil {
		return err
	}

	return nil
}


func (p productRepository) DecreaseStock(ctx context.Context, productId uuid.UUID, quantity int) error {
	err := p.dbGorm.WithContext(ctx).Model(&domain.Product{}).Where("id = ?", productId).Update("stock", gorm.Expr("stock - ?", quantity)).Error
	if err != nil {
		return err
	}

	return nil
}