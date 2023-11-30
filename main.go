package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// "github.com/i3ygene/test-project-api/src/router/router"
func main() {
	// r := Server()
	// r.Run
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/healthcheck", func(c *gin.Context) {
		// no content for healthcheck, only write http status
		c.Status(http.StatusOK)
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
	})

	router.Run(":8080")
}
