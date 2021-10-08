package repository

import (
	"test/model"
	"test/module/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetUserList(data map[string]interface{}) ([]*model.User, error) {
	var (
		err    error
		result = make([]*model.User, 0)
	)

	err = u.db.Find(&result, data).Error

	return result, err

}

func (u *UserRepository) GetUser(result *model.User) (*model.User, error) {
	var err error
	err = u.db.First(&result).Error

	return result, err
}

func (u *UserRepository) CreateUser(result *model.User) (*model.User, error) {
	var err error
	err = u.db.Create(&result).Error

	return result, err
}

func (u *UserRepository) UpdateUser(result *model.User) (*model.User, error) {
	var err error
	err = u.db.Save(&result).Error

	return result, err
}

func (u *UserRepository) ModifyUser(result *model.User, data map[string]interface{}) (*model.User, error) {
	var err error
	err = u.db.Model(&result).Updates(data).Error

	return result, err
}

func (u *UserRepository) DeleteUser(result *model.User) error {
	var err error
	err = u.db.Delete(&result).Error

	return err
}
