package repository

import (
	"context"
	"my-echo-chat_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}


func NewUser(db *mongo.Database) domain.UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}



func (u userRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	err := u.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, nil 
		}
		return domain.User{}, err
	}

	return user, nil
}


func (u userRepository) Insert(ctx context.Context, user domain.User) error {
	_, err := u.collection.InsertOne(ctx, user)
	return err
}

func (u userRepository) FindUser(ctx context.Context) (domain.User, error) {
	var user domain.User

	err := u.collection.FindOne(ctx, bson.M{}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, nil 
		}
		return domain.User{}, err
	}

	return user, nil
}

func (u userRepository) FindUserId(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}

	err = u.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, mongo.ErrNoDocuments
		}
		return domain.User{}, err
	}

	return user, nil
}