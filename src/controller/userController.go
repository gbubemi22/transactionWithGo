package controller

import (
	"github.com/gin-gonic/gin"
	//"github.com/gbubemi22/transaction/src/model"
	"github.com/gbubemi22/transaction/src/repository"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type UserController struct {
	userRepository *repository.UserRepository
}

// NewUserController creates a new instance of the UserController.
func NewUserController(userRepository *repository.UserRepository) *UserController {
	return &UserController{userRepository}
}

func (uc *UserController) GetOneUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user ID from the URL parameter.
		userIDStr := c.Param("id")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		// Retrieve the user from the repository using the user ID.
		user, err := uc.userRepository.GetUserByID(uint(userID))
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the user as a JSON response.
		c.JSON(http.StatusOK, user)
	}
}



func (ac *UserController) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement logic to retrieve users from the repository.
		users, err := ac.userRepository.GetAllUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the list of users as a JSON response.
		c.JSON(http.StatusOK, users)
	}
}