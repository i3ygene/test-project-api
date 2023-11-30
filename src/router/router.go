package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
	})

	router.Run(":8080")
}
