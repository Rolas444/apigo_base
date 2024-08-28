package services

import (
	"errors"

	"github.com/Rolas444/apigo_base/domain/models"
	"github.com/Rolas444/apigo_base/domain/repository"
	"github.com/Rolas444/apigo_base/utils"
)

type UserService interface {
	Login(credentials *models.Credentials) (string, error)
	Create(user *models.User) (*models.User, error)
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User, id uint) (*models.User, error)
	Delete(id uint) error
}

type UserServiceImpl struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
	}
}

func (s *UserServiceImpl) Login(credentials *models.Credentials) (string, error) {
	user, err := s.UserRepo.FindByEmail(credentials.Email)
	if err != nil {
		return "", errors.New("error interno")
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	compError := utils.CheckPasswordHash(credentials.Password, user.Password)
	if compError != nil {
		return "", errors.New("invalid password")
	}
	return utils.GenerateToken(user.ID)
}

func (s *UserServiceImpl) Create(user *models.User) (*models.User, error) {
	userExists, _ := s.UserRepo.FindByEmail(user.Email)
	if userExists != nil {
		return nil, errors.New("user already exists")
	}
	user.Password, _ = utils.HashPassword(user.Password)
	return s.UserRepo.Create(user)
}

func (s *UserServiceImpl) FindAll() ([]models.User, error) {
	return s.UserRepo.FindAll()
}

func (s *UserServiceImpl) FindByID(id uint) (*models.User, error) {
	return s.UserRepo.FindByID(id)
}

func (s *UserServiceImpl) FindByEmail(email string) (*models.User, error) {
	return s.UserRepo.FindByEmail(email)
}

func (s *UserServiceImpl) Update(user *models.User, id uint) (*models.User, error) {
	userToUpdate, err := s.UserRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user.Password == "" || user.Password == userToUpdate.Password {
		user.Password = userToUpdate.Password
	} else {
		user.Password, err = utils.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
	}

	return s.UserRepo.Update(user)
}

func (s *UserServiceImpl) Delete(id uint) error {
	return s.UserRepo.Delete(id)
}
