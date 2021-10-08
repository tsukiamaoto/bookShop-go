package user

import "test/model"

type Repository interface {
	GetUserList(map[string]interface{}) ([]*model.User, error)
	GetUser(result *model.User) (*model.User, error)
	CreateUser(result *model.User) (*model.User, error)
	UpdateUser(result *model.User) (*model.User, error)
	ModifyUser(result *model.User, data map[string]interface{}) (*model.User, error)
	DeleteUser(result *model.User) error
}
