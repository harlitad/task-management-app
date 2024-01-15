package model

type Task struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Status      string `gorm:"type:varchar(255)" json:"status"`
	UserId      int64  `gorm:"ForeignKey:fk_task_user"`
}
