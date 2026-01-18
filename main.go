package main

import (
	"auth-api/config"
	"auth-api/handlers"
	"auth-api/middleware"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", handlers.LoginHandler)
	r.POST("/register", handlers.RegisterHandler)
	r.GET("/me", middleware.JWTAuth(), func(c *gin.Context) {
		userIDAny, _ := c.Get("user_id")
		roleAny, _ := c.Get("role")

		userID := userIDAny.(int)
		role := roleAny.(string)

		c.JSON(200, gin.H{
			"user_id": userID,
			"role":    role,
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
