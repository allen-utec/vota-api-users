package domain

import (
	"gorm.io/gorm"
)

/* Entities */

type User struct {
	gorm.Model
	Nickname string
}

type UserRepository interface {
	Create(user User) (User, error)
	GetAll() ([]User, error)
}
