package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name       string     `json:"name"`
	MiddleName string     `json:"middle_name"`
	LastName   *string    `json:"last_name"`
	Email      *string    `json:"email" gorm:"unique"`
	BirthDate  *time.Time `json:"birth_date"`
	Phone      *string    `json:"phone"`
	Sex        *string    `json:"sex"`
	Address    *string    `json:"address"`
	Document   *string    `json:"document" gorm:"unique"`
	DoctypeID  *uint      `json:"doctype_id"`
	Doctype    Doctype    `json:"doctype" gorm:"foreignKey:DoctypeID"`
	Status     bool       `json:"status" gorm:"default:true"`
}
