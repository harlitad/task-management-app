package repository

import (
	"context"
	"fmt"

	"github.com/harlitad/task-management-app/internal/model"
	"github.com/harlitad/task-management-app/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	Client            *mongo.Client
	Logger            logger.ILogger
	DBName            string
	UserCollectionStr string
	UserCollection    *mongo.Collection
}

func NewUserMongoRepository(client *mongo.Client) IUserRepository {
	collection := client.Database("tma").Collection("users")
	return &UserMongoRepository{
		Client:            client,
		DBName:            "tma",
		UserCollectionStr: "users",
		UserCollection:    collection,
	}
}

func (r *UserMongoRepository) Create(user model.User) error {
	_, err := r.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserMongoRepository) GetByEmail(email string) (model.User, error) {
	var user model.User
	filter := bson.D{{Key: "email", Value: email}} // Correct usage of bson.E
	err := r.UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, fmt.Errorf("user by email not found") // No user found
		}
		return model.User{}, err
	}
	return user, nil
}
