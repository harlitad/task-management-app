package usecase

import (
	"fmt"

	"example.com/task-management-app/model"
	"example.com/task-management-app/service"
)

type ITaskUsecase interface {
	Get(id string, userId string) (model.Task, error)
	Delete(id string, userId string) error
	List(userId string) ([]model.Task, error)
	Create(newTask model.Task) (model.Task, error)
}

type TaskUsecase struct {
	TaskService service.ITaskService
}

func NewTaskUsecase(taskService service.ITaskService) ITaskUsecase {
	return &TaskUsecase{
		TaskService: taskService,
	}
}

func (t *TaskUsecase) Get(id string, userId string) (model.Task, error) {
	task, err := t.TaskService.Get(id, userId)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (t *TaskUsecase) Delete(id string, userId string) error {
	err := t.TaskService.Delete(id, userId)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskUsecase) List(userId string) ([]model.Task, error) {
	tasks, err := t.TaskService.List(userId)
	if err != nil {
		fmt.Println("error?")
		return nil, err
	}
	return tasks, nil
}

func (u *TaskUsecase) Create(task model.Task) (model.Task, error) {
	res, err := u.TaskService.Create(task)
	if err != nil {
		return model.Task{}, err
	}
	return res, nil
}
