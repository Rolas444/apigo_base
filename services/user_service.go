package services

import (
	"github.com/Rolas444/apigo_base/domain/models"
	"github.com/Rolas444/apigo_base/domain/repository"
)

type UserService interface {
	Create(user *models.User) (*models.User, error)
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
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

func (s *UserServiceImpl) Create(user *models.User) (*models.User, error) {
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

func (s *UserServiceImpl) Update(user *models.User) (*models.User, error) {
	return s.UserRepo.Update(user)
}

func (s *UserServiceImpl) Delete(id uint) error {
	return s.UserRepo.Delete(id)
}
