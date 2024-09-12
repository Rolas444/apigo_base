package routes

import (
	"github.com/Rolas444/apigo_base/controller"
	"github.com/Rolas444/apigo_base/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine, userController *controller.UserController) {
	userGroup := router.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())
	{
		// userGroup.POST("/login", userController.Login)
		userGroup.GET("/", userController.FindAllUsers)
		userGroup.POST("/", userController.CreateUser)
		userGroup.GET("/:id", userController.FindUserByID)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
		userGroup.POST("/logout", userController.Logout)
	}

}
