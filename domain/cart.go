package domain

import (
	"context"
	"my-golang-service-pos/dto"
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserId    string    `json:"user_id"`
	ProductId uuid.UUID `json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductId" json:"product"`
	Quantity  int       `json:"quantity"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}


type CartRepository interface {
	Insert(ctx context.Context, userId string, productId uuid.UUID, quantity int) error
	GetByCartId(ctx context.Context, userId string) ([]Cart , error)
	GetAll(ctx context.Context) ([]Cart, error)
	Delete(ctx context.Context, id string) error
	ClearCart(ctx context.Context, userId string) error
}	


type CartService interface {
	AddOrUpdate(ctx context.Context, req dto.AddCartReq) error
	GetCartByUserId(ctx context.Context, userId string) (dto.CartFullRes , error)
	GetAll(ctx context.Context) ([]dto.CartFullRes, error)
	DeleteCartById(ctx context.Context, id string) error
}






