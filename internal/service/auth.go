package service

import (
	"context"
	"errors"
	"log"
	"my-golang-service-pos/domain"
	"my-golang-service-pos/dto"
	"my-golang-service-pos/internal/config"
	"my-golang-service-pos/internal/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewUser(cnf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return &userService{
		conf:           cnf,
		userRepository: userRepository,
	}
}

func (d userService) Register(ctx context.Context, req dto.UserData) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	newUser := domain.User{
		Id:        uuid.New(),
		FullName:  req.Fullname,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	err = d.userRepository.Insert(ctx, newUser)

	if err != nil {
		return "", err
	}

	tokenStr, err := utils.GenerateToken(newUser.Id.String(), d.conf.Jwt)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (d userService) Login(ctx context.Context, req dto.UserRequest) (dto.UserResponse, error) {
	user, err := d.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.UserResponse{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.UserResponse{}, errors.New("authentication failed")
	}

	// Generate token dari utils
	tokenStr, err := utils.GenerateToken(user.Id.String(), d.conf.Jwt)
	if err != nil {
		return dto.UserResponse{}, errors.New("authentication failed (token error)")
	}

	return dto.UserResponse{
		Token: tokenStr,
	}, nil
}

func (d userService) GetProfile(ctx context.Context, id string) (domain.User, error) {

	return d.userRepository.FindUserId(ctx, id)
}

func (d userService) Logout(ctx context.Context, userId string) error {
	user, err := d.userRepository.FindUserId(ctx, userId)
	if err != nil {
		log.Printf("error logout: %v", err)
		return err
	}

	if user.Id.String() != userId {
		log.Printf("error logout: %v", err)
		return errors.New("you are not authorized")
	}
	return nil
}
