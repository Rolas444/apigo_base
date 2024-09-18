package repository

import "github.com/Rolas444/apigo_base/domain/models"

type PerosnRepository interface {
	Create(person *models.Person) (*models.Person, error)
	FindAll() ([]models.Person, error)
	FindByID(id uint) (*models.Person, error)
	Update(person *models.Person) (*models.Person, error)
	Delete(id uint) error
}
