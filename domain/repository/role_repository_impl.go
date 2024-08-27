package repository

import (
	"github.com/Rolas444/apigo_base/domain/models"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	DB *gorm.DB
}

func (r *RoleRepositoryImpl) Create(role *models.Role) (*models.Role, error) {
	err := r.DB.Create(role).Error
	return role, err
}

func (r *RoleRepositoryImpl) FindAll() ([]models.Role, error) {
	var roles []models.Role
	err := r.DB.Find(&roles).Error
	return roles, err
}

func (r *RoleRepositoryImpl) FindByID(id uint) (*models.Role, error) {
	var role models.Role
	err := r.DB.First(&role, id).Error
	return &role, err
}

func (r *RoleRepositoryImpl) FindByName(name string) (*models.Role, error) {
	var role models.Role
	err := r.DB.Where("name = ?", name).First(&role).Error
	return &role, err
}

func (r *RoleRepositoryImpl) Update(role *models.Role) (*models.Role, error) {
	err := r.DB.Save(role).Error
	return role, err
}

func (r *RoleRepositoryImpl) Delete(id uint) error {
	err := r.DB.Delete(&models.Role{}, id).Error
	return err
}
