package model

type User struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(255)" json:"name"`
	Email string `gorm:"type:varchar(255)" json:"email"`
}
