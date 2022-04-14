package users

import (
	"Belajar-Go-Echo/domains/users"
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.UserRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateUser(user users.User) error {
	response := r.DB.Create(&user)

	if response.RowsAffected < 1 {
		return fmt.Errorf("error insert user")
	}

	return nil
}

func (r *repository) GetAllUser() []users.User {
	user := []users.User{}
	r.DB.Find(&user)

	return user
}

func (r *repository) GetUserById(id int) (users.User, error) {
	user := users.User{}
	response := r.DB.Where("id = ?", id).Find(&user)

	if response.RowsAffected < 1 {
		return user, response.Error
	}

	return user, nil
}
