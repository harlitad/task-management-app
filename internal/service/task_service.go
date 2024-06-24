package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/harlitad/task-management-app/internal/model"
	"github.com/harlitad/task-management-app/internal/repository"
	"github.com/harlitad/task-management-app/pkg/logger"
)

type ITaskService interface {
	GetTasks(c context.Context) ([]model.Task, error)
	GetTasksToday(c context.Context) ([]model.Task, error)
	GetTaskById(c context.Context, id uuid.UUID) (model.Task, error)
	UpdateTaskById(c context.Context, id uuid.UUID, updatedTask model.Task) (model.Task, error)
	Delete(c context.Context, id uuid.UUID) error
	Get(c context.Context, id uuid.UUID, userId uuid.UUID) (model.Task, error)
	Create(c context.Context, newTask model.Task) (model.Task, error)
}

type TaskService struct {
	TaskRepository repository.ITaskRepository
	Logger         logger.ILogger
}

func NewTaskService(logger logger.ILogger, taskRepository repository.ITaskRepository) ITaskService {
	return &TaskService{
		TaskRepository: taskRepository,
		Logger:         logger,
	}
}

func (s *TaskService) GetTasks(c context.Context) ([]model.Task, error) {
	tasks, err := s.TaskRepository.GetTasks(c)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) GetTasksToday(c context.Context) ([]model.Task, error) {
	tasks, err := s.TaskRepository.GetTasksToday(c)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) GetTaskById(c context.Context, id uuid.UUID) (model.Task, error) {
	task, err := s.TaskRepository.GetTaskById(c, id)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (s *TaskService) Create(c context.Context, task model.Task) (model.Task, error) {
	task.Id = uuid.New()
	task.Status = model.TaskStatusTodo
	err := s.TaskRepository.Create(c, task)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (s *TaskService) UpdateTaskById(c context.Context, id uuid.UUID, updatedTask model.Task) (model.Task, error) {
	task, err := s.TaskRepository.UpdateTaskById(c, id, updatedTask)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (s *TaskService) Delete(c context.Context, id uuid.UUID) error {
	err := s.TaskRepository.Delete(c, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) Get(c context.Context, id uuid.UUID, userId uuid.UUID) (model.Task, error) {
	task, err := s.TaskRepository.Get(c, id, userId)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}
