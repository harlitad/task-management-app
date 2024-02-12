package service

import (
	"example.com/task-management-app/model"
	"example.com/task-management-app/repository"
	"example.com/task-management-app/utils"
	"github.com/google/uuid"
)

type IUserService interface {
	Create(user model.User) error
	GetByEmail(email string) (model.User, error)
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) Create(user model.User) error {
	hashingPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashingPassword
	user.Id = uuid.NewString()

	// calling to repository
	err = s.UserRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetByEmail(email string) (model.User, error) {
	user, err := s.UserRepository.GetByEmail(email)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}
