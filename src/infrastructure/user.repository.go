package infrastructure

import (
	"github.com/allen-utec/vota-api-users/src/domain"
)

type UserRepository struct{}

func (r *UserRepository) Create(user domain.User) (domain.User, error) {
	result := dbConn.Create(&user)

	return user, result.Error
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	result := dbConn.Find(&users)

	return users, result.Error
}
