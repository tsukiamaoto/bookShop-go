package service

import (
	"shopCart/model"
	service "shopCart/module/service/implement"
	repo "shopCart/module/repository"
)

type Users interface {
	GetUserList(map[string]interface{}) ([]*model.User, error)
	GetUser(in *model.User) (*model.User, error)
	CreateUser(in *model.User) (*model.User, error)
	UpdateUser(in *model.User) (*model.User, error)
	ModifyUser(in *model.User, data map[string]interface{}) (*model.User, error)
	DeleteUser(in *model.User) error
}

type Services struct {
	Users Users
}

type Deps struct {
	Repos *repo.Repositories
}

func NewServices(deps Deps) *Services {
	usersService := service.NewUserSerivce(deps.Repos.Users)

	return &Services{
		Users : usersService,
	}
}
