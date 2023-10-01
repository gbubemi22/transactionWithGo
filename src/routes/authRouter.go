package routes


import (
	"github.com/gin-gonic/gin"
	"github.com/gbubemi22/transaction/src/controller"
 )
 
 // Set up authentication-related routes.
func AuthRoutes(incomingRoutes *gin.Engine, authController *controller.AuthController) {
	// Create an "/api/v1/auth" route group.
	authGroup := incomingRoutes.Group("/api/v1/auth")
 
	authGroup.POST("/register", authController.Signup())
	authGroup.POST("/login", authController.Login())
 }


 