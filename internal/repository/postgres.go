package repository

import (
	"fmt"
	"log"

	"github.com/harlitad/task-management-app/config"
	"github.com/harlitad/task-management-app/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgreClient(config config.Config) (*gorm.DB, error) {
	// Construct DSN

	dsn := fmt.Sprintf("postgres://%s?user=%s&password=%s&dbname=%s&port=%s&sslmode=disable&TimeZone=Asia/Shanghai", config.PostgreHost, config.PostgreUsername, config.PostgrePassword, config.PostgreDBName, config.PostgrePort)
	dialect := postgres.Open(dsn)

	// Open database connection
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Configure database connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to configure database connection pooling: %w", err)
	}
	sqlDB.SetMaxOpenConns(100) // Set maximum open connections
	sqlDB.SetMaxIdleConns(10)  // Set maximum idle connections

	// Run migrations
	// if err := migrations.RunMigrations(dsn); err != nil {
	// 	return nil, fmt.Errorf("failed to run migrations: %w", err)
	// }

	err = db.AutoMigrate(&model.Task{}, &model.User{})
	if err != nil {
		log.Fatal("Failed to migrate " + err.Error())
	}

	return db, nil
}

// func runMigrations(dsn string) error {
// 	migrationPath := "../../migrations"

// 	// Split DSN to create a new connection string for golang-migrate
// 	m, err := migrate.New(
// 		"file://"+migrationPath,
// 		fmt.Sprintf("postgres://%s", dsn),
// 	)
// 	if err != nil {
// 		return fmt.Errorf("failed to create migrate instance: %w", err)
// 	}

// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 		return fmt.Errorf("failed to apply migrations: %w", err)
// 	}

// 	return nil
// }
