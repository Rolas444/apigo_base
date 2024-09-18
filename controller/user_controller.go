package controller

import (
	"strconv"

	"github.com/Rolas444/apigo_base/domain/models"
	"github.com/Rolas444/apigo_base/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (ctrl *UserController) Login(c *gin.Context) {
	var credentials models.Credentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": "Invalid credentials"})
		return
	}

	// println("Credentials: ", fmt.Sprintf("%v", &credentials))
	token, err := ctrl.UserService.Login(&credentials)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Set the token in a secure, HttpOnly cookie
	// c.SetCookie("token", token, 60*60*24, "/", "true", true, true)
	c.JSON(200, gin.H{"token": token, "message": "Login successful"})
}

func (ctrl *UserController) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "true", true, true)
	c.JSON(200, gin.H{"message": "Logout successful"})
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := ctrl.UserService.Create(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": createdUser})
}

func (ctrl *UserController) FindAllUsers(c *gin.Context) {
	users, err := ctrl.UserService.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": users})
}

func (ctrl *UserController) FindUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := ctrl.UserService.FindByID(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid ID"})
		return
	}

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := ctrl.UserService.Update(&user, uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": updatedUser})
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid ID"})
		return
	}

	err = ctrl.UserService.Delete(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted"})
}
