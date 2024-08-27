package repository

import (
	"github.com/Rolas444/apigo_base/domain/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (r *UserRepositoryImpl) Create(user *models.User) (*models.User, error) {
	err := r.DB.Create(user).Error
	return user, err
}

func (r *UserRepositoryImpl) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user *models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) Update(user *models.User) (*models.User, error) {
	err := r.DB.Save(user).Error
	return user, err
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	err := r.DB.Delete(&models.User{}, id).Error
	return err
}
