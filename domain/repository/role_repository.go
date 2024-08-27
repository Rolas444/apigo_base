package repository

import "github.com/Rolas444/apigo_base/domain/models"

type RoleRepository interface {
	FindAll() ([]models.Role, error)
	FindByID(id uint) (*models.Role, error)
	FindByName(name string) (*models.Role, error)
	Create(role *models.Role) (*models.Role, error)
	Update(role *models.Role) (*models.Role, error)
	Delete(id uint) error
}
