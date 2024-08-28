package main

import (
	"log"

	"github.com/Rolas444/apigo_base/config"
	"github.com/Rolas444/apigo_base/controller"
	"github.com/Rolas444/apigo_base/domain/repository"
	"github.com/Rolas444/apigo_base/middleware"
	"github.com/Rolas444/apigo_base/routes"
	"github.com/Rolas444/apigo_base/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.Connect()

	config.MigrateTables(config.DB)

	roleRepo := &repository.RoleRepositoryImpl{DB: config.DB}
	userRepo := &repository.UserRepositoryImpl{DB: config.DB}
	initService := &services.InitService{RoleRepo: roleRepo, UserRepo: userRepo}
	initController := &controller.InitController{InitService: initService}
	userService := services.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	initController.Initialize()
	router := gin.Default()
	//use CORS
	router.Use(middleware.CORSMiddleware())

	routes.UserRoute(router, &userController)

	//public routes
	router.GET("/login", userController.Login)

	router.Run(":8080")

}
