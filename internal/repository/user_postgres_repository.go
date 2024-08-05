package repository

import (
	"github.com/harlitad/task-management-app/internal/model"
	"gorm.io/gorm"
)

type UserPostgresRepository struct {
	PostgreClient *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) IUserRepository {
	return &UserPostgresRepository{
		PostgreClient: db,
	}
}

func (r *UserPostgresRepository) Create(user model.User) error {
	err := r.PostgreClient.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserPostgresRepository) GetByEmail(email string) (model.User, error) {
	user := model.User{}
	err := r.PostgreClient.Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
