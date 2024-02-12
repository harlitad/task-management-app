package repository

import (
	"fmt"

	"example.com/task-management-app/model"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	List(userId string) ([]model.Task, error)
	Get(id string, userId string) (model.Task, error)
	Delete(id string, userId string) error
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

func (r *TaskRepository) List(userId string) ([]model.Task, error) {
	tasks := []model.Task{}
	err := r.PostgreClient.Model(&model.Task{}).Where("user_id = ?", userId).Scan(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) Get(id string, userId string) (model.Task, error) {
	task := model.Task{}
	err := r.PostgreClient.Model(&task).Where("id = ?", id).Where("user_id = ?", userId).First(&task).Error
	if err != nil {
		fmt.Println(err)
		return task, err
	}
	return task, nil
}

func (r *TaskRepository) Delete(id string, userId string) error {
	err := r.PostgreClient.Where("id = ?", id).Where("user_id = ?", userId).Delete(&model.Task{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) Create(newTask model.Task) error {
	err := r.PostgreClient.Create(&newTask).Error
	if err != nil {
		return err
	}
	return nil
}
