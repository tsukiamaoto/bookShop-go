package implement

import (
	"github.com/tsukiamaoto/bookShop-go/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetUserList() ([]*model.User, error) {
	var (
		users = make([]*model.User, 0)
	)

	err := u.db.Model(&model.User{}).Find(&users).Error

	return users, err

}

func (u *UserRepository) GetUser(user *model.User) (*model.User, error) {
	err := u.db.First(&user).Error

	return user, err
}

func (u *UserRepository) GetUserById(userId uint) (*model.User, error) {
	var user *model.User
	err := u.db.First(&user).Error

	return user, err
}

func (u *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	err := u.db.Create(&user).Error

	return user, err
}

func (u *UserRepository) UpdateUser(user *model.User) (*model.User, error) {
	err := u.db.Model(&model.User{ID: user.ID}).Updates(&user).Error

	return user, err
}

func (u *UserRepository) DeleteUser(userId uint) error {
	err := u.db.Delete(&model.User{ID: userId}).Error

	return err
}
