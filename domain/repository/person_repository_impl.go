package repository

import (
	"github.com/Rolas444/apigo_base/domain/models"
	"gorm.io/gorm"
)

type PersonRepository struct {
	DB *gorm.DB
}

func (r *PersonRepository) Create(person *models.Person) (*models.Person, error) {
	err := r.DB.Create(person).Error
	return person, err
}

func (r *PersonRepository) FindAll() ([]models.Person, error) {
	var persons []models.Person
	err := r.DB.Find(&persons).Error
	return persons, err
}

func (r *PersonRepository) FindByID(id uint) (*models.Person, error) {
	var person models.Person
	err := r.DB.First(&person, id).Error
	return &person, err
}

func (r *PersonRepository) Update(person *models.Person) (*models.Person, error) {
	err := r.DB.Save(person).Error
	return person, err
}

func (r *PersonRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Person{}, id).Error
	return err
}
