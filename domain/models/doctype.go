package models

import "gorm.io/gorm"

type Doctype struct {
	gorm.Model
	Name   string `json:"name"`
	Status bool   `json:"status" gorm:"default:true"`
}
