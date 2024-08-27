package controller

import (
	"os"
	"time"

	"github.com/Rolas444/apigo_base/config"
	"github.com/Rolas444/apigo_base/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Invalid parameters"})
		return
	}
	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost) //10
	if err != nil {
		c.JSON(500, gin.H{"error": "Error hashing password"})
		return
	}
	//create user
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(200, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Invalid parameters"})
		return
	}

	var user models.User
	config.DB.Where("email = ?", body.Email).First(&user)
	if user.ID == 0 {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}

	//Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}) //30 days
	tokenString, err := token.SignedString(os.Getenv("SECRET"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Error generating token"})
	}

	c.JSON(200, gin.H{"token": tokenString, "message": "User logged in"})
}

func GetUser(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func GetUserById(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.JSON(200, &user)
}
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}
func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}
