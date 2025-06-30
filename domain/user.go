package domain

import (
	"context"
	"github.com/google/uuid"

	"time"
)

type User struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"fullName"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
	Insert(ctx context.Context, user User) error
	FindUser(ctx context.Context) (User, error)
	FindUserId(ctx context.Context, id string) (User, error)
}
