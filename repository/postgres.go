package repository

import (
	"fmt"
	"log"

	"example.com/task-management-app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreClient(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.PostgreHost, config.PostgreUsername, config.PostgrePassword, config.PostgreDBName, config.PostgrePort)

	dialect := postgres.Open(dsn)

	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// err = db.AutoMigrate(&Task{}, &User{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return db, nil
}
