package repository

import (
	"fmt"
	"log"

	"example.com/task-management-app/config"
	"example.com/task-management-app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgreClient(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.PostgreHost, config.PostgreUsername, config.PostgrePassword, config.PostgreDBName, config.PostgrePort)

	dialect := postgres.Open(dsn)

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = db.AutoMigrate(&model.Task{}, &model.User{})
	if err != nil {
		log.Fatal("Failed to migrate " + err.Error())
	}
	return db, nil
}
