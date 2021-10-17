package service

import (
	"shopCart/model"
	"shopCart/module/user"
)

type UserService struct {
	repo user.Repository
}

func NewUserSerivce(repo user.Repository) user.Service {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUserList(data map[string]interface{}) ([]*model.User, error) {
	return u.repo.GetUserList(data)
}

func (u *UserService) GetUser(result *model.User) (*model.User, error) {
	return u.repo.GetUser(result)
}

func (u *UserService) CreateUser(result *model.User) (*model.User, error) {
	return u.repo.CreateUser(result)
}

func (u *UserService) UpdateUser(result *model.User) (*model.User, error) {
	return u.repo.UpdateUser(result)
}

func (u *UserService) ModifyUser(result *model.User, data map[string]interface{}) (*model.User, error) {
	return u.repo.ModifyUser(result, data)
}

func (u *UserService) DeleteUser(result *model.User) error {
	return u.repo.DeleteUser(result)
}
