package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id" bson:"_id,omitempty"`
	Email     string `json:"email"`
	FullName  string `json:"fullname"`
	Password  string `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
	Insert(ctx context.Context, user User) error
	FindUser(ctx context.Context) (User, error)
	FindUserId(ctx context.Context, id string) (User, error)
}


