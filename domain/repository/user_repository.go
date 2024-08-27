package repository

import "github.com/Rolas444/apigo_base/domain/models"

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id uint) error
}
