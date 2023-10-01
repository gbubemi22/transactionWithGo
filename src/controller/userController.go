package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "net/http"
    "strconv"
    "github.com/gbubemi22/transaction/src/model"
)

type UserController struct {
    db *gorm.DB
}

// NewUserController creates a new instance of the UserController.
func NewUserController(db *gorm.DB) *UserController {
    return &UserController{db}
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

        // Retrieve the user from the database using the user ID.
        var user model.User
        if err := uc.db.First(&user, userID).Error; err != nil {
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

func (uc *UserController) GetUsers() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Implement logic to retrieve users from the database.
        var users []model.User
        if err := uc.db.Find(&users).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Return the list of users as a JSON response.
        c.JSON(http.StatusOK, users)
    }
}
