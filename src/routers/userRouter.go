package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gbubemi22/transaction/src/controller"
)

// Set up user-related routes.
func UserRoutes(incomingRoutes *gin.Engine, userController *controller.UserController) {
	// Create an "/api/v1/users" route group.
	userGroup := incomingRoutes.Group("/api/v1/users")

	// Use the handler functions directly without calling them.
	userGroup.GET("/:id", userController.GetOneUser())
	userGroup.GET("/", userController.GetUsers())
}
