package config

import (
	"log"

	"github.com/Rolas444/apigo_base/domain/models"
	"gorm.io/gorm"
)

func MigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(&models.Role{}, &models.User{}, &models.Doctype{}, &models.ValidToken{})
	if err != nil {
		log.Fatalf("Error migrating tables: %v", err)
	}
}
