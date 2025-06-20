package repository

import (
	"context"
	"my-golang-service-pos/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db : db,
	}
}

func (u userRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}

	return user, nil
}

func (u userRepository) Insert(ctx context.Context, user domain.User) error {

	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return err
}

func (u userRepository) FindUser(ctx context.Context) (domain.User, error) {
	var user domain.User

	err := u.db.WithContext(ctx).First(&user).Error
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
	err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, mongo.ErrNoDocuments
		}
		return domain.User{}, err
	}

	return user, nil
}
