package auth

import (
	"Belajar-Go-Echo/domains/auth"
	"Belajar-Go-Echo/domains/users"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.AuthRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Login(credential auth.Auth) error {
	response := r.DB.Where("email = ? AND password = ?", credential.Email, credential.Password).Find(&users.User{})

	if response.RowsAffected < 1 {
		return response.Error
	}

	return nil
}