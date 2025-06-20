package domain

import (
	"context"
	"my-golang-service-pos/dto"
)

type AuthService interface {
	Register(ctx context.Context, req dto.UserData) (string, error)
	Login(ctx context.Context, req dto.UserRequest) (dto.UserResponse, error)
	GetProfile(ctx context.Context, id string) (User, error)
	Logout(ctx context.Context, id string) error
}

type AuthServiceTesting interface {
	TestRegister(ctx context.Context, req dto.UserData) (string, error)
}
