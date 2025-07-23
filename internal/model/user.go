package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeleteAt  time.Time `json:"deleteAt,omitempty"`
}
