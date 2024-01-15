package repository

import "gorm.io/gorm"

type UserRepository struct {
	PostgreClient *gorm.DB
}

func (u *UserRepository) Get() {

}
