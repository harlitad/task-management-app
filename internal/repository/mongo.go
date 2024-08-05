package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/harlitad/task-management-app/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(config config.Config) (*mongo.Client, error) {
	if err := validateConfig(config); err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		return nil, err
	}
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin",
		config.MongoDBUsername,
		config.MongoDBPassword,
		config.MongoDBHost,
		config.MongoDBPort,
		config.MongoDB,
	)

	clientOptions := options.Client().ApplyURI(uri)
	timeout := 10 * time.Second
	clientOptions.ConnectTimeout = &timeout

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return client, nil
}

func validateConfig(config config.Config) error {
	if config.MongoDBUsername == "" || config.MongoDBPassword == "" ||
		config.MongoDBHost == "" || config.MongoDBPort == 0 || config.MongoDB == "" {
		return fmt.Errorf("invalid MongoDB configuration")
	}
	return nil
}
