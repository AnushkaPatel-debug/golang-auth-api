package handlers

import (
	"golang-auth-api/database"
	"golang-auth-api/models"
	"golang-auth-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)
	if user.Role == "" {
		user.Role = "user"
	}

	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

func Login(c *gin.Context) {

	var input models.User
	var user models.User

	c.BindJSON(&input)
	database.DB.Where("email= ?", input.Email).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return

	}
	token, err := utils.GenerateJWT(user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func Profile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to profile route",
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}
