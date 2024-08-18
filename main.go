package main

import (
	"github.com/Rolas444/apigo_base/config"
	"github.com/Rolas444/apigo_base/initializers"
	"github.com/Rolas444/apigo_base/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	initializers.LoadEnvVariables()
	config.Connect()
	initializers.SyncDatabase()
}

func main() {
	router := gin.New()
	//config.Connect()
	Init()
	routes.UserRoute(router)
	router.Run(":8080")

}
