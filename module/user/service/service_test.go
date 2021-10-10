package service_test

import (
	"errors"
	"testing"

	mocks "test/mocks/module/user"
	"test/model"
	"test/module/user/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserList(t *testing.T) {
	mockRepository := new(mocks.Repository)
	mockUsers := []*model.User{
		&model.User{
			ID:       uint(1),
			Username: "123",
			Password: "123",
		},
		&model.User{
			ID:       uint(2),
			Username: "123",
			Password: "123",
		},
	}

	t.Run("Success", func(t *testing.T) {
		mockRepository.
			On("GetUserList", mock.Anything).
			Return(mockUsers, nil).Once()

		userRepo := service.NewUserSerivce(mockRepository)
		users, err := userRepo.GetUserList(make(map[string]interface{}))

		assert.NoError(t, err)
		assert.NotNil(t, users)

		// check mockRepository to call "On" function
		mockRepository.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {
		mockRepository.
			On("GetUserList", mock.Anything).
			Return(nil, errors.New("Got error")).Once()

		userRepo := service.NewUserSerivce(mockRepository)
		users, err := userRepo.GetUserList(make(map[string]interface{}))

		assert.Error(t, err)
		assert.Nil(t, users)

		// check mockRepository to call "On" function
		mockRepository.AssertExpectations(t)
	})
}

func TestGetUser(t *testing.T) {
	mockRepository := new(mocks.Repository)
	mockUser := &model.User{
		ID:       uint(1),
		Username: "123",
		Password: "123",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepository.
			On("GetUser", mock.Anything).
			Return(mockUser, nil).Once()

		userRepo := service.NewUserSerivce(mockRepository)
		user, err := userRepo.GetUser(&model.User{ ID: uint(1)})

		assert.NoError(t, err)
		assert.NotNil(t, user)

		// check mockRepository to call "On" function
		mockRepository.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {
		mockRepository.
			On("GetUser", mock.Anything).
			Return(nil, errors.New("Got error")).Once()

		userRepo := service.NewUserSerivce(mockRepository)
		user, err := userRepo.GetUser(&model.User{ ID: uint(1)})

		assert.Error(t, err)
		assert.Nil(t, user)

		// check mockRepository to call "On" function
		mockRepository.AssertExpectations(t)
	})
}

func TestCreateUser(t *testing.T) {
	mockRepository := new(mocks.Repository)
	mockUser := &model.User{
		ID:       uint(1),
		Username: "123",
		Password: "123",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepository.
			On("CreateUser", mock.Anything).
			Return(mockUser, nil).Once()

		userRepo := service.NewUserSerivce(mockRepository)
		user, err := userRepo.CreateUser(mockUser)

		assert.NoError(t, err)
		assert.NotNil(t, user)

		// check mockRepository to call "On" function
		mockRepository.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {
		mockRepository.
			On("CreateUser", mock.Anything).
			Return(nil, errors.New("Got error")).Once()

		userRepo := service.NewUserSerivce(mockRepository)
		user, err := userRepo.CreateUser(mockUser)

		assert.Error(t, err)
		assert.Nil(t, user)

		// check mockRepository to call "On" function
		mockRepository.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	mockRepository := new(mocks.Repository)
	mockUser := &model.User{
		ID:       uint(1),
		Username: "123",
		Password: "123",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepository.
			On("UpdateUser", mock.Anything).
			Return(mockUser, nil).Once()

		userRepo := service.NewUserSerivce(mockRepository)
		user, err := userRepo.UpdateUser(mockUser)

		assert.NoError(t, err)
		assert.NotNil(t, user)

		// check mockRepository to call "On" function
		mockRepository.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {
		mockRepository.
			On("UpdateUser", mock.Anything).
			Return(nil, errors.New("Got error")).Once()

		userRepo := service.NewUserSerivce(mockRepository)
		user, err := userRepo.UpdateUser(mockUser)

		assert.Error(t, err)
		assert.Nil(t, user)

		// check mockRepository to call "On" function
		mockRepository.AssertExpectations(t)
	})
}

func TestModifyUser(t *testing.T) {
	mockRepository := new(mocks.Repository)
	mockUsers := []*model.User{
		&model.User{
			ID:       uint(1),
			Username: "123",
			Password: "123",
		},
		&model.User{
			ID:       uint(2),
			Username: "123",
			Password: "123",
		},
	}

	newUser := map[string]interface{} {
		"username": "456",
		"password": "456",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepository.
			On("ModifyUser", mock.Anything, mock.Anything).
			Return(mockUsers[0], nil).Once()

		userRepo := service.NewUserSerivce(mockRepository)
		_, err := userRepo.ModifyUser(&model.User{ID: 1}, newUser)

		assert.NoError(t, err)

		// check mockRepository to call "On" function
		mockRepository.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {
		mockRepository.
			On("ModifyUser", mock.Anything, mock.Anything).
			Return(nil, errors.New("Got error")).Once()

		userRepo := service.NewUserSerivce(mockRepository)
		_, err := userRepo.ModifyUser(&model.User{ID: 1}, newUser)

		assert.Error(t, err)

		// check mockRepository to call "On" function
		mockRepository.AssertExpectations(t)
	})
}

func TestDeleteUser (t *testing.T) {
	mockRepository := new(mocks.Repository)

	t.Run("Success", func(t *testing.T) {
		mockRepository.
			On("DeleteUser", mock.Anything).
			Return(nil).Once()

		userRepo := service.NewUserSerivce(mockRepository)
		err := userRepo.DeleteUser(&model.User{ID: 1})

		assert.NoError(t, err)

		// check mockRepository to call "On" function
		mockRepository.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {
		mockRepository.
			On("DeleteUser", mock.Anything).
			Return(errors.New("Got error")).Once()

		userRepo := service.NewUserSerivce(mockRepository)
	  err := userRepo.DeleteUser(&model.User{ID: 1})

		assert.Error(t, err)

		// check mockRepository to call "On" function
		mockRepository.AssertExpectations(t)
	})
}