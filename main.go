package main

import (
	"golang-auth-api/database"
	"golang-auth-api/handlers"
	"golang-auth-api/middleware"
	"golang-auth-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{})
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server Running",
		})

	})
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)
	r.GET("/profile", middleware.AutheMiddleware(), handlers.Profile)
	r.GET("/users", middleware.AutheMiddleware(), middleware.AdminMiddleware(), handlers.GetUsers)
	r.Run(":8090")
}
