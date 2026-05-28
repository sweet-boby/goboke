package repository

import (
	"errors"
	"goboke/internal/model"
	"goboke/internal/util/password"
	"time"
)

type MemoryUserRepository struct {
	users  []model.User
	nextID int
}

func NewMemoryUserRepository() *MemoryUserRepository {
	// In-memory storage
	pd := "12345678"
	hash, _ := password.HashPassword(pd)
	var users = []model.User{
		{ID: 1, Username: "lzw", Phone: "12345678", Password: "12345678", PasswordHash: hash, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	var nextID = 2

	return &MemoryUserRepository{
		users:  users,
		nextID: nextID,
	}
}

func (r *MemoryUserRepository) FindUserByUsername(username string) *model.User {

	for i, item := range r.users {
		if item.Username == username {
			return &r.users[i]
		}
	}

	return nil
}

func (r *MemoryUserRepository) FindUserByPhone(phone string) *model.User {
	for i, item := range r.users {
		if item.Phone == phone {
			return &r.users[i]
		}
	}

	return nil
}

func (r *MemoryUserRepository) FindUserByID(id int) *model.User {
	for i, item := range r.users {
		if item.ID == id {
			return &r.users[i]
		}
	}
	return nil
}

func (r *MemoryUserRepository) Create(user model.User) (*model.User, error) {
	// TODO: Check if username already exists
	// TODO: Check if email already exists
	for _, item := range r.users {
		if item.Username == user.Username {
			return nil, errors.New("username exists")
		}
	}
	// TODO: Hash password
	hash, err := password.HashPassword(user.Password)
	if err != nil {

		return nil, errors.New("hash password fail")
	}
	user.ID = r.nextID
	r.nextID += 1

	user.PasswordHash = hash

	// TODO: Create user and add to users slice
	r.users = append(r.users, user)

	return &user, nil
}

func (r *MemoryUserRepository) Update(id int, user model.User) (*model.User, error) {
	return nil, nil
}

func (r *MemoryUserRepository) Delete(id int) error {
	return nil
}
