package usecase

import (
	"github.com/google/uuid"
	"github.com/harlitad/task-management-app/internal/model"
	"github.com/harlitad/task-management-app/internal/service"
	"github.com/harlitad/task-management-app/pkg/logger"
	"golang.org/x/net/context"
)

type ITaskUsecase interface {
	GetTasks(c context.Context) ([]model.Task, error)
	GetTasksToday(c context.Context) ([]model.Task, error)
	Create(c context.Context, newTask model.Task) (model.Task, error)
	GetTaskById(c context.Context, id uuid.UUID) (model.Task, error)
	UpdateTask(c context.Context, updateTask model.Task) (model.Task, error)
	Delete(c context.Context, id uuid.UUID) error
}

type TaskUsecase struct {
	TaskService service.ITaskService
	Logger      logger.ILogger
}

func NewTaskUsecase(logger logger.ILogger, taskService service.ITaskService) ITaskUsecase {
	return &TaskUsecase{
		TaskService: taskService,
		Logger:      logger,
	}
}

func (t *TaskUsecase) UpdateTask(c context.Context, updateTask model.Task) (model.Task, error) {
	task, err := t.TaskService.UpdateTaskById(c, updateTask.Id, updateTask)
	if err != nil {
		t.Logger.Errorf("error update task, %v", err)
		return model.Task{}, err
	}
	return task, nil
}

func (t *TaskUsecase) GetTaskById(c context.Context, id uuid.UUID) (model.Task, error) {
	task, err := t.TaskService.GetTaskById(c, id)
	if err != nil {
		t.Logger.Errorf("error get task by id, %v", err)
		return model.Task{}, err
	}
	return task, nil
}

func (t *TaskUsecase) GetTasks(c context.Context) ([]model.Task, error) {
	tasks, err := t.TaskService.GetTasks(c)
	if err != nil {
		t.Logger.Errorf("error get tasks, %v", err)
		return nil, err
	}
	return tasks, nil
}

func (t *TaskUsecase) GetTasksToday(c context.Context) ([]model.Task, error) {
	tasks, err := t.TaskService.GetTasksToday(c)
	if err != nil {
		t.Logger.Errorf("error get tasks, %v", err)
		return nil, err
	}
	return tasks, nil
}

func (u *TaskUsecase) Create(c context.Context, newTask model.Task) (model.Task, error) {

	task, err := u.TaskService.Create(c, newTask)
	if err != nil {
		u.Logger.Errorf("error create task, %v", err)
		return model.Task{}, err
	}

	return task, nil
}

func (t *TaskUsecase) Delete(c context.Context, id uuid.UUID) error {
	err := t.TaskService.Delete(c, id)
	if err != nil {
		t.Logger.Errorf("error delete task, %v", err)
		return err
	}
	return nil
}

func (t *TaskUsecase) Get(c context.Context, id uuid.UUID, userId uuid.UUID) (model.Task, error) {
	task, err := t.TaskService.Get(c, id, userId)
	if err != nil {
		t.Logger.Errorf("error get task, %v", err)
		return model.Task{}, err
	}
	return task, nil
}
