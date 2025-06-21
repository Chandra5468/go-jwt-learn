package routes

import (
	"github.com/Chandra5468/go-jwt-learn/controllers"
	"github.com/Chandra5468/go-jwt-learn/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:userId", controllers.GetUser())
}
