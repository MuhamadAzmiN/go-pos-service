package domain

import (
	"context"
	"my-echo-chat_service/dto"
)
type AuthService interface {
	Register(ctx context.Context, req dto.UserData)  (string, error) 
	Login(ctx context.Context, req dto.UserRequest) (dto.UserResponse, error)
	GetProfile(ctx context.Context, id string) (User, error)
	Logout(ctx context.Context, id string) error
}

