package repository

import (
	"example.com/task-management-app/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user model.User) error
	GetByEmail(email string) (model.User, error)
}

type UserRepository struct {
	PostgreClient *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		PostgreClient: db,
	}
}

func (r *UserRepository) Create(user model.User) error {
	err := r.PostgreClient.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByEmail(email string) (model.User, error) {
	user := model.User{}
	err := r.PostgreClient.Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
