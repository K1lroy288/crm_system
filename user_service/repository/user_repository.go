package repository

import (
	model "user-service/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *model.User) (bool, error) {
	var exist model.User
	res := r.DB.Where("username = ?", user.Username).First(&exist).Error
	if res != nil {
		return res == nil, r.DB.Create(user).Error
	}
	return res == nil, nil
}

func (r *UserRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *UserRepository) GetUserByLastname(lastname string) (model.User, error) {
	var user model.User
	err := r.DB.Where("last_name = ?", lastname).First(&user).Error
	return user, err
}
