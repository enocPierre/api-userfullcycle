package database

import (
	"github.com/goexpert/apis/internal/entity"
)

type UserInterface interface {
	Creatte(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type PriductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	delete(id string) error
}
