package repository

import "goboke/internal/model"

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindUserByUsername(username string) *model.User
	FindUserByPhone(phone string) *model.User
	FindUserByID(id int) *model.User
	Create(user model.User) (*model.User, error)
	Update(id int, user model.User) (*model.User, error)
	Delete(id int) error
}
