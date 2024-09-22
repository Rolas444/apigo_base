package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Name         string     `json:"name"`
	DepartmentID uint       `json:"department_id"`
	Department   Department `json:"department" gorm:"foreignKey:DepartmentID"`
	MemberID     uint       `json:"member_id"`
	Member       Member     `json:"member" gorm:"foreignKey:MemberID"`
	Status       bool       `json:"status" gorm:"default:true"`
}
