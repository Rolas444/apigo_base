package routes

import (
	"github.com/Rolas444/apigo_base/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/signup", controller.SignUp)
	router.GET("/", controller.GetUser)
	router.POST("/", controller.CreateUser)
	router.GET("/:id", controller.GetUserById)
	router.PUT("/:id", controller.UpdateUser)
	router.DELETE("/:id", controller.DeleteUser)
}
