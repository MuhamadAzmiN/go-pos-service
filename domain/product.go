package domain

import (
	"context"
	"github.com/google/uuid"
	"my-golang-service-pos/dto"
)

type Product struct {
	Id    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name  string    `json:"name" gorm:"column:name"`
	Sku   string    `json:"sku" gorm:"column:sku"`
	Price int   `json:"price" gorm:"column:price"`
	Stock int       `json:"stock" gorm:"column:stock"`
}

type ProductRepository interface {
	FindAll(ctx context.Context) ([]Product, error)
	FindById(ctx context.Context, id string) (Product, error)
	Insert(ctx context.Context, product Product) error
	Delete(ctx context.Context, id string) error
	DecreaseStock(ctx context.Context, productId uuid.UUID, quantity int) error
}

type ProductService interface {
	GetProductList(ctx context.Context) ([]dto.ProductResponse, error)
	GetProductById(ctx context.Context, id string) (dto.ProductData, error)
	CreateProduct(ctx context.Context, req dto.ProductRequest) error
	DeleteProduct(ctx context.Context, id string) error
}
