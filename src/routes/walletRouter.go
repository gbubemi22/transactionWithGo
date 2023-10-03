package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gbubemi22/transaction/src/controller"
)

// WalletRoutes sets up wallet-related routes.
func WalletRoutes(incomingRoutes *gin.Engine, walletController *controller.WalletController) {
	// Create a "/api/v1/wallet" route group.
	walletGroup := incomingRoutes.Group("/api/v1/wallet")

	// Define a route to create a wallet.
	walletGroup.POST("/:userID", walletController.Createwallet())

	
}
