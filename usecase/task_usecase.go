package usecase

import (
	"fmt"

	"example.com/task-management-app/model"
	"example.com/task-management-app/service"
)

type ITaskUsecase interface {
	Get(id int64) (model.Task, error)
	List() ([]model.Task, error)
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

func (t *TaskUsecase) Get(id int64) (model.Task, error) {
	task, err := t.TaskService.Get(id)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (t *TaskUsecase) List() ([]model.Task, error) {
	tasks, err := t.TaskService.List()
	if err != nil {
		fmt.Println("error?")
		return nil, err
	}
	return tasks, nil
}

func (u *TaskUsecase) Create(newTask model.Task) (model.Task, error) {
	err := u.TaskService.Create(newTask)
	if err != nil {
		return newTask, err
	}
	return newTask, nil
}
