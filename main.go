package main

import (
	"auth-api/config"
	"auth-api/handlers"
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

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
