package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/harlitad/task-management-app/internal/dto"
	"gorm.io/gorm"
)

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in-progress"
	TaskStatusReview     TaskStatus = "review"
	TaskStatusDone       TaskStatus = "done"
)

var TaskStatusMap = map[TaskStatus]string{
	TaskStatusDone:       "1",
	TaskStatusInProgress: "2",
	TaskStatusReview:     "3",
	TaskStatusTodo:       "4",
}

type Task struct {
	Id          uuid.UUID      `gorm:"type:uuid;primaryKey;not null" bson:"id" json:"id"`
	Title       string         `gorm:"type:varchar(255)" bson:"title" json:"title"`
	Description string         `gorm:"type:text" bson:"description" json:"description"`
	Status      TaskStatus     `gorm:"type:varchar(20)" bson:"status" json:"status"`
	AssigneeId  uuid.UUID      `gorm:"type:uuid" bson:"assignee_id" json:"assignee_id"`
	CreatorId   uuid.UUID      `gorm:"type:uuid" bson:"creator_id" json:"creator_id"`
	DueDate     time.Time      `gorm:"type:timestamp with time zone" bson:"due_date" json:"due_date"`
	CreatedAt   time.Time      `gorm:"type:timestamp with time zone" bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"type:timestamp with time zone" bson:"updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" bson:"deleted_at" json:"deleted_at"`
}

func MapCreateTaskRequest(req dto.CreateTaskRequest) Task {
	var dueDate time.Time
	if req.DueDate != "" {
		dueDate, _ = time.Parse(time.RFC3339, req.DueDate)
	}
	parsedUUID, _ := uuid.Parse(req.AssigneeId)
	return Task{
		Title:       req.Title,
		Description: req.Description,
		AssigneeId:  parsedUUID,
		DueDate:     dueDate,
		CreatedAt:   time.Now().Local().UTC(),
		UpdatedAt:   time.Now().Local().UTC(),
	}
}

func MapTaskRequest(req dto.CreateTaskRequest) Task {
	var dueDate time.Time
	if req.DueDate != "" {
		dueDate, _ = time.Parse(time.RFC3339, req.DueDate)
	}
	parsedUUID, _ := uuid.Parse(req.AssigneeId)
	return Task{
		Title:       req.Title,
		Description: req.Description,
		AssigneeId:  parsedUUID,
		DueDate:     dueDate,
		CreatedAt:   time.Now().Local().UTC(),
		UpdatedAt:   time.Now().Local().UTC(),
	}
}

func MapUpdateTaskRequest(req dto.UpdateTaskRequest) Task {
	var dueDate time.Time
	if req.DueDate != "" {
		dueDate, _ = time.Parse(time.RFC3339, req.DueDate)
	}
	parsedUUID, _ := uuid.Parse(req.AssigneeId)
	return Task{
		Title:       req.Title,
		Description: req.Description,
		AssigneeId:  parsedUUID,
		DueDate:     dueDate,
		Status:      TaskStatus(req.Status),
	}
}

func MapToTaskResponse(t Task) dto.Task {
	return dto.Task{
		Id:          t.Id.String(),
		Title:       t.Title,
		Description: t.Description,
		Status:      string(t.Status),
		CreatorId:   t.CreatorId.String(),
		AssigneeId:  t.AssigneeId.String(),
		DueDate:     t.DueDate,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
