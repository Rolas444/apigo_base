package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status" gorm:"default:true"`
}
