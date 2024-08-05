package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/harlitad/task-management-app/internal/model"
)

type ITaskRepository interface {
	GetTasks(c context.Context) ([]model.Task, error)
	GetTasksToday(c context.Context) ([]model.Task, error)
	GetTaskById(c context.Context, id uuid.UUID) (model.Task, error)
	Create(c context.Context, newTask model.Task) error
	UpdateTaskById(c context.Context, id uuid.UUID, newTask model.Task) (model.Task, error)
	Delete(c context.Context, id uuid.UUID) error
	Get(c context.Context, id uuid.UUID, userId uuid.UUID) (model.Task, error)
}

type IUserRepository interface {
	Create(user model.User) error
	GetByEmail(email string) (model.User, error)
}
