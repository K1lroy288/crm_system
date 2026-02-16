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

func (s *UserService) GetUserByLastname(lastname string) (*model.MasterDTO, error) {
	user, err := s.repo.GetUserByLastname(lastname)
	if err != nil {
		return nil, err
	}

	userDTO := &model.MasterDTO{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return userDTO, nil
}

func (s *UserService) GetMasters() ([]model.MasterDTO, error) {
	masters, err := s.repo.GetMasters()
	if err != nil {
		return nil, err
	}

	var mastersDTO []model.MasterDTO
	for _, master := range masters {
		masterDTO := &model.MasterDTO{
			ID:        master.ID,
			Username:  master.Username,
			FirstName: master.FirstName,
			LastName:  master.LastName,
		}

		mastersDTO = append(mastersDTO, *masterDTO)
	}

	return mastersDTO, nil
}

func (s *UserService) GetMastersByIDs(mastersIDs []uint) ([]model.MasterDTO, error) {
	masters, err := s.repo.GetMastersByIDs(mastersIDs)
	if err != nil {
		return nil, err
	}

	var mastersDTO []model.MasterDTO
	for _, master := range masters {
		masterDTO := &model.MasterDTO{
			ID:        master.ID,
			Username:  master.Username,
			FirstName: master.FirstName,
			LastName:  master.LastName,
		}

		mastersDTO = append(mastersDTO, *masterDTO)
	}

	return mastersDTO, nil
}
