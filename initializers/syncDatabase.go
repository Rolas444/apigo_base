package initializers

import (
	"github.com/Rolas444/apigo_base/config"
	"github.com/Rolas444/apigo_base/models"
)

func SyncDatabase() {
	config.DB.AutoMigrate(&models.Role{}, &models.User{})
}
