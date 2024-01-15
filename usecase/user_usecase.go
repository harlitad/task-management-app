package usecase

import (
	"example.com/task-management-app/model"
	"example.com/task-management-app/service"
)

type IUserUsecase interface {
	RegisterNewUser(user model.User)
}

type UserUsecase struct {
	TaskService service.ITaskService
}

func NewUserUsecase(taskService service.ITaskService) ITaskUsecase {
	return &TaskUsecase{
		TaskService: taskService,
	}
}

func (u *TaskUsecase) Register(user model.User) (model.User, error) {
	return user, nil
}
