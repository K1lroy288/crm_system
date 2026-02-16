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

func (r *UserRepository) GetMasters() ([]model.User, error) {
	var masters []model.User
	err := r.DB.Raw(`
		SELECT u.* FROM users u
		INNER JOIN user_roles ur ON u.id = ur.user_id
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE r.role_name = ? AND u.deleted_at IS NULL
	`, "master").Scan(&masters).Error

	return masters, err
}

func (r *UserRepository) GetMastersByIDs(masterIDs []uint) ([]model.User, error) {
	var masters []model.User
	err := r.DB.Raw(`
		SELECT u.* FROM users u
		JOIN user_roles ur ON u.id = ur.user_id
		JOIN roles r ON ur.role_id = r.id
		WHERE r.role_name = ? AND u.deleted_at IS NULL 
		AND u.id IN ?
	`, "master", masterIDs).Scan(&masters).Error

	return masters, err
}
