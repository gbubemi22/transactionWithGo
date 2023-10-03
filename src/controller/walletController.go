package controller


import (
	"github.com/jinzhu/gorm"
    "net/http"
    "strconv"
	"github.com/gin-gonic/gin"
	"github.com/gbubemi22/transaction/src/model"
)



type WalletController struct {
	db *gorm.DB
 }
 
//  type ErrorResponse struct {
// 	Error string `json:"error"`
//  }
 

 func NewWalletController(db *gorm.DB) *WalletController {
	return &WalletController{db}
 }


 func (wc *WalletController) Createwallet() gin.HandlerFunc {
	return func(c *gin.Context) {

       // Get the UserID from the params
	  userIDStr := c.Param("userID")
	  userID, err := strconv.Atoi(userIDStr)
	  if err != nil {
		 c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid userID"})
		 return
	  }

	  // Find the user by ID
	  var user model.User
	  if err := wc.db.First(&user, userID).Error; err != nil {
		 c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		 return
	  }

	  // Create the wallet with the UserID and FirstName as UserName
	  wallet := model.Wallet{
		 UserID:   uint(userID),
		 UserName: user.FirstName,
		 // You can set other wallet properties as needed
	  }

	  // Save the wallet to the database
	  if err := wc.db.Create(&wallet).Error; err != nil {
		 c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create wallet"})
		 return
	  }

	  c.JSON(http.StatusCreated, wallet)



	}


}