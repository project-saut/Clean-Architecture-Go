package users_test

import (
	"Belajar-Go-Echo/domains/users"
	user "Belajar-Go-Echo/domains/users"
	"Belajar-Go-Echo/domains/users/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

/*
func CreateUser(email, password string) *user.User {
	DB := database.ConnectDB()
	u := user.User{
		Email:    email,
		Password: password,
	}
	if err := DB.Create(&u).Error; err != nil {
		return nil
	}

	return &u
}
*/
func TestCreate(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUserData := users.User{
		Email:    "sautmanurungaja28@gmail.com",
		Password: "sautaja23",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("CreateUser", mock.Anything).Return(nil).Once()

		userHandler := user.UserRepository(mockUserRepo)
		err := userHandler.CreateUser(mockUserData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("CreateUser", mock.Anything).Return(errors.New("error unit test")).Once()

		userHandler := user.UserRepository(mockUserRepo)
		err := userHandler.CreateUser(mockUserData)

		assert.Error(t, err)
	})
}

func TestGetAllUser(t *testing.T) {

}
