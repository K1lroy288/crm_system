package repository

import (
	"time"
	model "user-service/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	err := r.DB.Where("username = ?", username).Preload("Roles").First(&user).Error
	return user, err
}

func (r *UserRepository) GetUserByLastname(lastname string) (model.User, error) {
	var user model.User
	err := r.DB.Where("last_name = ?", lastname).First(&user).Error
	return user, err
}

func (r *UserRepository) GetUsersByRole(role string) ([]model.User, error) {
	var users []model.User
	err := r.DB.Raw(`
		SELECT u.* FROM users u
		INNER JOIN user_roles ur ON u.id = ur.user_id
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE r.role_name = ? AND u.deleted_at IS NULL
	`, role).Scan(&users).Error

	return users, err
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

func (r *UserRepository) GetUserInfo(id uint) (model.User, error) {
	var user model.User
	err := r.DB.Model(&model.User{ID: id}).Scan(&user).Error
	return user, err
}

func (r *UserRepository) UpdateUser(user *model.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) GetUserById(id uint) (*model.User, error) {
	var user *model.User
	err := r.DB.Model(&model.User{ID: id}).Scan(&user).Error
	return user, err
}

func (r *UserRepository) GetRoles() ([]model.Role, error) {
	var roles []model.Role
	err := r.DB.Table("roles").Find(&roles).Error
	return roles, err
}

func (r *UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	err := r.DB.
		Table("users").
		Preload("Roles").
		Where("deleted_at IS NULL").
		Find(&users).
		Error

	return users, err
}

func (r *UserRepository) DeleteUser(id uint) error {
	return r.DB.Model(&model.User{ID: id}).Update("deleted_at", time.Now()).Error
}
