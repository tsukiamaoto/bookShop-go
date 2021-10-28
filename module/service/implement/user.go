package implement

import (
	"shopCart/model"
	repo "shopCart/module/repository"
)

type UsersService struct {
	repo repo.Users
}

func NewUserSerivce(repo repo.Users) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (u *UsersService) GetUserList(data map[string]interface{}) ([]*model.User, error) {
	return u.repo.GetUserList(data)
}

func (u *UsersService) GetUser(result *model.User) (*model.User, error) {
	return u.repo.GetUser(result)
}

func (u *UsersService) CreateUser(result *model.User) (*model.User, error) {
	return u.repo.CreateUser(result)
}

func (u *UsersService) UpdateUser(result *model.User) (*model.User, error) {
	return u.repo.UpdateUser(result)
}

func (u *UsersService) ModifyUser(result *model.User, data map[string]interface{}) (*model.User, error) {
	return u.repo.ModifyUser(result, data)
}

func (u *UsersService) DeleteUser(result *model.User) error {
	return u.repo.DeleteUser(result)
}
