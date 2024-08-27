package controller

import (
	"log"

	"github.com/Rolas444/apigo_base/services"
)

type InitController struct {
	InitService *services.InitService
}

func (ctrl *InitController) Initialize() {
	err := ctrl.InitService.InitRolesAndAdmin()
	if err != nil {
		// c.JSON(500, gin.H{"error": err.Error()})
		// return
		log.Fatalf("Error initializing schema: %v", err)
	}
	// c.JSON(http.StatusOK, gin.H{"message": "Initialization Successfull"})
}
