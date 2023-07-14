package database

import "github.com/goexpert/apis/internal/entity"

type UserInterface interface {
	Creatte(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
