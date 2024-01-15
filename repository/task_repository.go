package repository

import (
	"example.com/task-management-app/model"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	List() ([]model.Task, error)
	Get(id int64) (model.Task, error)
	Create(newTask model.Task) error
}

type TaskRepository struct {
	PostgreClient *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{
		PostgreClient: db,
	}
}

func (r *TaskRepository) List() ([]model.Task, error) {
	tasks := []model.Task{}
	err := r.PostgreClient.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) Get(id int64) (model.Task, error) {
	task := model.Task{}
	err := r.PostgreClient.First(&task, id).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *TaskRepository) Create(newTask model.Task) error {
	err := r.PostgreClient.Create(&newTask).Error
	if err != nil {
		return err
	}
	return nil
}
