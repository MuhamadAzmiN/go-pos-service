package repository

import (
	"context"
	"errors"
	"my-golang-service-pos/domain"
	"my-golang-service-pos/internal/iface"
	"strings"

	"gorm.io/gorm"
)

type userRepository struct {
	dbGorm iface.IGorm
	db     iface.ISqlx
}

func NewUser(dbGorm iface.IGorm, db iface.ISqlx) domain.UserRepository {
	return &userRepository{
		dbGorm: dbGorm,
		db:     db,
	}
}


func (u userRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	err := u.dbGorm.Model(&domain.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, nil
		}

		return domain.User{}, err
	}
	return user, nil

}

func (u userRepository) Insert(ctx context.Context, user domain.User) error {
	err := u.dbGorm.Create(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	return nil
}

func (u userRepository) FindUser(ctx context.Context) (domain.User, error) {
	var user domain.User

	err := u.dbGorm.WithContext(ctx).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}

	return user, nil
}

func (u userRepository) FindUserId(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	err := u.dbGorm.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}

	return user, nil
}
