package implement

import (
	"shopCart/model"
	repo "shopCart/module/repository"
)

type UsersService struct {
	repo repo.Users
}

func NewUsersService(repo repo.Users) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (u *UsersService) GetUserList() ([]*model.User, error) {
	return u.repo.GetUserList()
}

func (u *UsersService) GetUser(user *model.User) (*model.User, error) {
	return u.repo.GetUser(user)
}

func (u *UsersService) GetUserById(userId uint) (*model.User, error) {
	return u.repo.GetUserById(userId)
}

func (u *UsersService) CreateUser(user *model.User) (*model.User, error) {
	return u.repo.CreateUser(user)
}

func (u *UsersService) UpdateUser(user *model.User) (*model.User, error) {
	return u.repo.UpdateUser(user)
}

func (u *UsersService) DeleteUser(userId uint) error {
	return u.repo.DeleteUser(userId)
}
