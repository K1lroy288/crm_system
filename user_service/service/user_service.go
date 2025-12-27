package service

import (
	"user-service/model"
	"user-service/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByUsername(username string) (model.User, error) {
	return s.repo.GetUserByUsername(username)
}

func (s *UserService) CreateUser(user *model.User) (bool, error) {
	return s.repo.CreateUser(user)
}
