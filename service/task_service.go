package service

import (
	"example.com/task-management-app/model"
	"example.com/task-management-app/repository"
)

type ITaskService interface {
	List() ([]model.Task, error)
	Get(id int64) (model.Task, error)
	Create(newTask model.Task) error
}

type TaskService struct {
	TaskRepository repository.ITaskRepository
}

func NewTaskService(taskRepository repository.ITaskRepository) ITaskService {
	return &TaskService{
		TaskRepository: taskRepository,
	}
}

func (s *TaskService) List() ([]model.Task, error) {
	tasks, err := s.TaskRepository.List()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) Get(id int64) (model.Task, error) {
	task, err := s.TaskRepository.Get(id)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (s *TaskService) Create(newTask model.Task) error {
	err := s.TaskRepository.Create(newTask)
	if err != nil {
		return err
	}
	return nil
}
