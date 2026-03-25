package service

import (
	"strconv"
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

func (s *UserService) CreateUser(userDTO model.UserDTO) error {
	roles, err := s.repo.GetRoles()
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	var currRoles []model.Role
	if len(userDTO.Roles) == 0 {
		for _, r := range roles {
			if r.RoleName == "user" {
				currRoles = append(currRoles, r)
				break
			}
		}
	} else {
		for _, cr := range userDTO.Roles {
			for _, r := range roles {
				if cr.ID == r.ID {
					currRoles = append(currRoles, r)
					break
				}
			}
		}
	}

	user := &model.User{
		Username:     userDTO.Username,
		FirstName:    userDTO.FirstName,
		LastName:     userDTO.Lastname,
		Surname:      userDTO.Surname,
		PasswordHash: hashedPassword,
		Roles:        currRoles,
	}

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
		Lastname:  user.LastName,
		Surname:   user.Surname,
	}
	return userDTO, nil
}

func (s *UserService) GetUsersByRole(role string) ([]model.UserDTO, error) {
	users, err := s.repo.GetUsersByRole(role)
	if err != nil {
		return nil, err
	}

	var usersDTO []model.UserDTO
	for _, user := range users {
		userDTO := &model.UserDTO{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			Lastname:  user.LastName,
			Surname:   user.Surname,
		}

		usersDTO = append(usersDTO, *userDTO)
	}

	return usersDTO, nil
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
			Lastname:  master.LastName,
			Surname:   master.Surname,
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
		Lastname:  user.LastName,
		Surname:   user.Surname,
	}

	return userDTO, err
}

func (s *UserService) UpdateUser(userDTO model.UserDTO) error {
	existingUser, err := s.repo.GetUserById(userDTO.ID)
	if err != nil {
		return err
	}

	existingUser.Username = userDTO.Username
	existingUser.FirstName = userDTO.FirstName
	existingUser.LastName = userDTO.Lastname
	existingUser.Surname = userDTO.Surname

	if userDTO.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		existingUser.PasswordHash = hashedPassword
	}

	roles, err := s.repo.GetRoles()
	if err != nil {
		return err
	}

	var currRoles []model.Role
	if len(userDTO.Roles) == 0 {
		for _, r := range roles {
			if r.RoleName == "user" {
				currRoles = append(currRoles, r)
				break
			}
		}
	} else {
		for _, cr := range userDTO.Roles {
			for _, r := range roles {
				if cr.ID == r.ID {
					currRoles = append(currRoles, r)
					break
				}
			}
		}
	}
	existingUser.Roles = currRoles

	return s.repo.UpdateUser(existingUser)
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

	err = s.repo.UpdateUser(user)
	return err
}

func (s *UserService) GetRoles() ([]model.RoleDTO, error) {
	roles, err := s.repo.GetRoles()
	if err != nil {
		return nil, err
	}

	var rolesDTO []model.RoleDTO
	for _, r := range roles {
		roleDTO := &model.RoleDTO{
			ID:   r.ID,
			Name: r.RoleName,
		}

		rolesDTO = append(rolesDTO, *roleDTO)
	}

	return rolesDTO, nil
}

func (s *UserService) GetUsers() ([]model.UserDTO, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	var usersDTO []model.UserDTO
	for _, u := range users {
		var roles []model.RoleDTO
		for _, r := range u.Roles {
			roleDTO := &model.RoleDTO{
				ID:   r.ID,
				Name: r.RoleName,
			}

			roles = append(roles, *roleDTO)
		}

		userDTO := &model.UserDTO{
			ID:        u.ID,
			Username:  u.Username,
			FirstName: u.FirstName,
			Lastname:  u.LastName,
			Surname:   u.Surname,
			Roles:     roles,
		}

		usersDTO = append(usersDTO, *userDTO)
	}

	return usersDTO, nil
}

func (s *UserService) DeleteUser(idString string) error {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	return s.repo.DeleteUser(uint(id))
}
