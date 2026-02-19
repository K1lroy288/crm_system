package service

import (
	"user-service/model"
	"user-service/repository"

	"golang.org/x/crypto/bcrypt"
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

func (s *UserService) GetUserByLastname(lastname string) (*model.UserDTO, error) {
	user, err := s.repo.GetUserByLastname(lastname)
	if err != nil {
		return nil, err
	}

	userDTO := &model.UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return userDTO, nil
}

func (s *UserService) GetMasters() ([]model.UserDTO, error) {
	masters, err := s.repo.GetMasters()
	if err != nil {
		return nil, err
	}

	var mastersDTO []model.UserDTO
	for _, master := range masters {
		masterDTO := &model.UserDTO{
			ID:        master.ID,
			Username:  master.Username,
			FirstName: master.FirstName,
			LastName:  master.LastName,
		}

		mastersDTO = append(mastersDTO, *masterDTO)
	}

	return mastersDTO, nil
}

func (s *UserService) GetMastersByIDs(mastersIDs []uint) ([]model.UserDTO, error) {
	masters, err := s.repo.GetMastersByIDs(mastersIDs)
	if err != nil {
		return nil, err
	}

	var mastersDTO []model.UserDTO
	for _, master := range masters {
		masterDTO := &model.UserDTO{
			ID:        master.ID,
			Username:  master.Username,
			FirstName: master.FirstName,
			LastName:  master.LastName,
		}

		mastersDTO = append(mastersDTO, *masterDTO)
	}

	return mastersDTO, nil
}

func (s *UserService) GetUserInfo(id uint) (*model.UserDTO, error) {
	user, err := s.repo.GetUserInfo(id)
	if err != nil {
		return nil, err
	}

	userDTO := &model.UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return userDTO, err
}

func (s *UserService) UpdateUserInfo(userDTO model.UserDTO) error {
	user, err := s.repo.GetUserById(userDTO.ID)
	if err != nil {
		return err
	}

	user.FirstName = userDTO.FirstName
	user.LastName = userDTO.LastName
	user.Username = userDTO.Username

	return s.repo.UpdateUserInfo(user)
}

func (s *UserService) ChangePassword(id uint, passDTO model.PasswordDTO) error {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(passDTO.CurrPass)); err != nil {
		return err
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(passDTO.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = newHashedPassword

	err = s.repo.UpdateUserInfo(user)
	return err
}
