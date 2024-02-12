package model

type User struct {
	Id       string `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationResponse struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

type UserAuthenticated struct {
	Name  string
	Id    string
	Email string
}
