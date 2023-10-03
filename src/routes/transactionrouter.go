package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gbubemi22/transaction/src/controller"
)



// WalletRoutes sets up wallet-related routes.
func TransactionRoutes(incomingRoutes *gin.Engine, transactionController *controller.TransactionController) {
	// Create a "/api/v1/wallet" route group.
	transactionGroup := incomingRoutes.Group("/api/v1/transactions")

	// Define a route to create a wallet.
	transactionGroup.POST("/transfer", transactionController.Transfer())

	
}