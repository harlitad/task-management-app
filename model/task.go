package model

type Task struct {
	Id          string `gorm:"type:varchar(255);primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Status      string `gorm:"type:varchar(255)" json:"status"`
	UserId      string `gorm:"type:varchar(255)" json:"user_id"`
}

type CreateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreateTaskResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
