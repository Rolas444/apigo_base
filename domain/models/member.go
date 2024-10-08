package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	PersonID      uint       `json:"person_id"`
	Person        Person     `json:"person" gorm:"foreignKey:PersonID"`
	BaptismDate   *time.Time `json:"baptism_date"`
	OriginChurch  *string    `json:"origin_church"`
	Participation bool       `json:"participation"`
	Status        bool       `json:"status" gorm:"default:true"`
}
