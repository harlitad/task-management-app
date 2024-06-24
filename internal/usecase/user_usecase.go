package usecase

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/harlitad/task-management-app/internal/model"
	"github.com/harlitad/task-management-app/internal/service"
	"github.com/harlitad/task-management-app/pkg/auth"
	"github.com/harlitad/task-management-app/pkg/utils"
)

type IUserUsecase interface {
	Create(user model.User) (model.User, error)
	Authentication(email string, password string) (string, model.User, error)
}

type UserUsecase struct {
	UserService service.IUserService
}

func NewUserUsecase(userService service.IUserService) IUserUsecase {
	return &UserUsecase{
		UserService: userService,
	}
}

func (u *UserUsecase) Create(user model.User) (model.User, error) {
	user, err := u.UserService.Create(user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserUsecase) Authentication(email string, password string) (string, model.User, error) {

	// get user by email
	user, err := u.UserService.GetByEmail(email)
	if err != nil {
		return "", model.User{}, err
	}

	// verify password
	check, err := utils.VerifyPassword(password, user.Password)
	if err != nil && !check {
		return "", model.User{}, err
	}

	// generate jwt
	claims := jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
		"id":    user.Id,
	}
	token, err := auth.GenerateToken(claims)

	return token, user, nil
}
