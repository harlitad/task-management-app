package usecase

import (
	"example.com/task-management-app/model"
	"example.com/task-management-app/pkg/auth"
	"example.com/task-management-app/service"
	"example.com/task-management-app/utils"
	"github.com/dgrijalva/jwt-go"
)

type IUserUsecase interface {
	Create(user model.User) error
	Authentication(user model.User) (string, error)
}

type UserUsecase struct {
	UserService service.IUserService
}

func NewUserUsecase(userService service.IUserService) IUserUsecase {
	return &UserUsecase{
		UserService: userService,
	}
}

func (u *UserUsecase) Create(user model.User) error {
	err := u.UserService.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) Authentication(reqUser model.User) (string, error) {

	// get user by email
	user, err := u.UserService.GetByEmail(reqUser.Email)
	if err != nil {
		return "", err
	}

	// verify password
	check, err := utils.VerifyPassword(reqUser.Password, user.Password)
	if err != nil && !check {
		return "", err
	}

	// generate jwt
	claims := jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
		"id":    user.Id,
	}
	token, err := auth.GenerateToken(claims)

	return token, nil
}
