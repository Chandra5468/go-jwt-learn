package main

import (
	"os"

	"github.com/Chandra5468/go-jwt-learn/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8082"
	}

	router := gin.New() // For router we are using gin

	router.Use(gin.Logger())

	// We will have two kind of routes authROuter, userRouter

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(ctx *gin.Context) { // No need to pass w responseWriter and r httpRequest
		ctx.JSON(200, gin.H{"success": "Access granted for api-1"}) // Send status and struct or object
	})

	router.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
