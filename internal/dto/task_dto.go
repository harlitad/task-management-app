package dto

import (
	"time"

	validator "github.com/AccelByte/justice-input-validation-go"
)

type Task struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatorId   string    `json:"creator_id"`
	AssigneeId  string    `json:"assignee_id"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

type Tasks struct {
	Data []Task `json:"data"`
}

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required~Title name is blank"`
	Description string `json:"description"`
	AssigneeId  string `json:"assignee_id" valid:"uuidv4"`
	DueDate     string `json:"due_date,omitempty" valid:"dateTime"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AssigneeId  string `json:"assignee_id" valid:"uuidv4"`
	DueDate     string `json:"due_date" valid:"dateTime"`
	Status      string `json:"status"`
}

type CreateTaskResponse struct {
	Task
}

func (a *CreateTaskRequest) Validate() error {
	_, err := validator.ValidateStruct(a)
	return err
}

func (a *UpdateTaskRequest) Validate() error {
	_, err := validator.ValidateStruct(a)
	return err
}
