package model

import "time"

type User struct {
	Id        string    `gorm:"type:uuid;primaryKey;not null" bson:"_id" json:"id"`
	Name      string    `gorm:"type:varchar(255)" bson:"name" json:"name"`
	Email     string    `gorm:"type:varchar(255);unique;not null" bson:"email" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" bson:"password" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" bson:"updated_at" json:"updated_at"`
}

type UserTask struct {
	Task Task
	User User
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

type UserAuthenticated struct {
	Name  string
	Id    string
	Email string
}
