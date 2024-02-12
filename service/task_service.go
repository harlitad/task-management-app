package service

import (
	"example.com/task-management-app/model"
	"example.com/task-management-app/repository"
	"github.com/google/uuid"
)

type ITaskService interface {
	List(userId string) ([]model.Task, error)
	Get(id string, userId string) (model.Task, error)
	Delete(id string, userId string) error
	Create(newTask model.Task) (model.Task, error)
}

type TaskService struct {
	TaskRepository repository.ITaskRepository
}

func NewTaskService(taskRepository repository.ITaskRepository) ITaskService {
	return &TaskService{
		TaskRepository: taskRepository,
	}
}

func (s *TaskService) List(userId string) ([]model.Task, error) {
	tasks, err := s.TaskRepository.List(userId)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) Get(id string, userId string) (model.Task, error) {
	task, err := s.TaskRepository.Get(id, userId)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (s *TaskService) Create(task model.Task) (model.Task, error) {
	task.Id = uuid.NewString()
	err := s.TaskRepository.Create(task)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (s *TaskService) Delete(id string, userId string) error {
	err := s.TaskRepository.Delete(id, userId)
	if err != nil {
		return err
	}
	return nil
}
