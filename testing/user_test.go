package testing_test

import (
	"Belajar-Go-Echo/domains/users"
	user "Belajar-Go-Echo/domains/users"
	"Belajar-Go-Echo/testing/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUserData := users.User{
		Email:    "sautmanurungaja28@gmail.com",
		Password: "sautaja23",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("CreateUser", mock.Anything).Return(nil).Once()

		userRepository := user.UserRepository(mockUserRepo)
		err := userRepository.CreateUser(mockUserData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("CreateUser", mock.Anything).Return(errors.New("error unit test")).Once()

		userRepository := user.UserRepository(mockUserRepo)
		err := userRepository.CreateUser(mockUserData)

		assert.Error(t, err)
	})
}

func TestGetAllUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUserData := []users.User{
		{
			Email:    "test1@gmail.com",
			Password: "test1",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetAllUser").Return(mockUserData).Once()

		userRepository := user.UserRepository(mockUserRepo)
		allUser := userRepository.GetAllUser()

		assert.Equal(t, mockUserData, allUser)
	})

	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("GetAllUser", mock.Anything).Return(mockUserData, errors.New("error unit test")).Once()

		userRepository := user.UserRepository(mockUserRepo)
		allUser := userRepository.GetAllUser()

		assert.Equal(t, mockUserData, allUser)
	})
}

func TestGetUserById(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUserData := users.User{
		Email:    "sautmanurungaja28@gmail.com",
		Password: "sautaja123",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetUserById", mock.Anything).Return(mockUserData, nil).Once()

		userRepository := user.UserRepository(mockUserRepo)
		userById, err := userRepository.GetUserById(1)

		assert.NoError(t, err)
		assert.Equal(t, mockUserData, userById)
	})

	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("GetUserById", mock.Anything).Return(mockUserData, errors.New("error unit test")).Once()

		userRepository := user.UserRepository(mockUserRepo)
		userById, err := userRepository.GetUserById(1)

		assert.Error(t, err)
		assert.Equal(t, mockUserData, userById)
	})
}
