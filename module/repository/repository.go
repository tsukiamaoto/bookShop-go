package repository

import (
	"shopCart/model"
	repo "shopCart/module/repository/implement"

	"gorm.io/gorm"
)

type Users interface {
	GetUserList(map[string]interface{}) ([]*model.User, error)
	GetUser(result *model.User) (*model.User, error)
	CreateUser(result *model.User) (*model.User, error)
	UpdateUser(result *model.User) (*model.User, error)
	ModifyUser(result *model.User, data map[string]interface{}) (*model.User, error)
	DeleteUser(result *model.User) error
}

type Repositories struct {
	Users Users
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users: repo.NewUserRepository(db),
	}
}