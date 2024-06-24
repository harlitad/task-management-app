package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/harlitad/task-management-app/internal/model"
	"github.com/harlitad/task-management-app/pkg/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

type TaskRepository struct {
	PostgreClient *gorm.DB
	Logger        logger.ILogger
}

func NewTaskRepository(logger logger.ILogger, db *gorm.DB) ITaskRepository {
	return &TaskRepository{
		PostgreClient: db,
		Logger:        logger,
	}
}

func (r *TaskRepository) UpdateTaskById(c context.Context, id uuid.UUID, newTask model.Task) (model.Task, error) {
	task := model.Task{
		Id: id,
	}
	res := r.PostgreClient.Model(&task).Clauses(clause.Returning{}).Updates(newTask)
	if res.Error != nil {
		return task, fmt.Errorf("failed update task, %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return task, model.ErrorRecordNotFound
	}
	return task, nil
}

func (r *TaskRepository) GetTasks(c context.Context) ([]model.Task, error) {
	tasks := []model.Task{}
	err := r.PostgreClient.Model(&model.Task{}).Scan(&tasks).Error
	if err != nil {
		return nil, fmt.Errorf("failed get task list from database, %w", err)
	}
	return tasks, nil
}

func (r *TaskRepository) GetTasksToday(c context.Context) ([]model.Task, error) {
	tasks := []model.Task{}
	err := r.PostgreClient.Model(&model.Task{}).Where("DATE(due_date) = ?", time.Now().Format("2006-01-02")).Scan(&tasks).Error
	if err != nil {
		return nil, fmt.Errorf("failed get task list from database, %w", err)
	}
	return tasks, nil
}

func (r *TaskRepository) GetTaskById(c context.Context, id uuid.UUID) (model.Task, error) {
	task := model.Task{}
	err := r.PostgreClient.Model(&task).Where("id = ?", id).First(&task).Error
	if err == gorm.ErrRecordNotFound {
		return task, model.ErrorRecordNotFound
	}
	if err != nil {
		return task, fmt.Errorf("failed get task by id from database, %w", err)
	}
	return task, nil
}

func (r *TaskRepository) Get(c context.Context, id uuid.UUID, userId uuid.UUID) (model.Task, error) {
	task := model.Task{}
	err := r.PostgreClient.Model(&task).Where("id = ?", id).Where("user_id = ?", userId).First(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *TaskRepository) Delete(c context.Context, id uuid.UUID) error {
	task := model.Task{
		Id: id,
	}
	del := r.PostgreClient.Delete(&task)
	if del.RowsAffected == 0 {
		return model.ErrorRecordNotFound
	}
	if del.Error != nil {
		return del.Error
	}
	return nil
}

func (r *TaskRepository) Create(c context.Context, newTask model.Task) error {
	err := r.PostgreClient.Create(&newTask).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) ListByAssigneeId(assigneeId string) ([]model.Task, error) {
	var tasks []model.Task
	err := r.PostgreClient.Model(&model.Task{}).Where("assignee_id = ?", assigneeId).Scan(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func (r *TaskRepository) GetTaskByAssigneeId(id uuid.UUID, assigneeId uuid.UUID) (model.Task, error) {
	var task model.Task
	err := r.PostgreClient.Model(&model.Task{}).Where(&model.Task{Id: id, AssigneeId: assigneeId}).First(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *TaskRepository) GetTaskByDueDate(date time.Time) ([]model.Task, error) {
	var tasks []model.Task
	err := r.PostgreClient.Model(&model.Task{}).Where(&model.Task{DueDate: date}).Scan(&tasks).Error
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *TaskRepository) UpdateStatus(id uuid.UUID, status string) error {
	err := r.PostgreClient.Where(&model.Task{Id: id}).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}
