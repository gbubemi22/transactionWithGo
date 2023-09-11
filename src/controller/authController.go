package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gbubemi22/transaction/src/model"
	"github.com/gbubemi22/transaction/src/repository"
	"github.com/gbubemi22/transaction/src/utils"
	"github.com/jinzhu/gorm"
	//"context"
	"net/http"
	//"time"
	"fmt"
 )

 type AuthController struct {
	userRepository *repository.UserRepository 
 }
 type ErrorResponse struct {
	Error string `json:"error"`
 }


 // NewAuthController creates a new instance of the AuthController.
func NewAuthController(userRepository *repository.UserRepository) *AuthController {
	return &AuthController{userRepository}
 }

 // Register handles user registration.

func (ac *AuthController) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
	
 
	   
	      var userData model.User
        if err := c.ShouldBindJSON(&userData); err != nil {
            c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON data: " + err.Error()})
            return
        }
       

        fmt.Println(userData)
	   
 
	  // Check if the email already exists in the database.
      existingUser, err := ac.userRepository.GetUserByEmail(userData.Email)
      if err != nil && !gorm.IsRecordNotFoundError(err) {
       c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
       }


	  fmt.Println(existingUser)
      if existingUser != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
         return
          }

 
	    // Check if the phone number already exists in the database.
	    existingUser, err = ac.userRepository.GetUserByPhone(userData.Phone)
	    if err != nil && !gorm.IsRecordNotFoundError(err) {
		   c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		   return
	    }
	    if existingUser != nil {
		   c.JSON(http.StatusConflict, gin.H{"error": "Phone number already exists"})
		   return
	    }
 
	    // Validate the user's password using the correct function name.existingUser
	    if err := utils.ValidatePasswordString(userData.Password); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid password: " + err.Error()})
		return
	 }
 
	 
	    // Hash the user's password before storing it.
	    hashedPassword, err := utils.HashPassword(userData.Password)
	    if err != nil {
		   c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to hash password"})
		   return
	    }

	     
 
	    userData.Password = hashedPassword
 
	    if err := ac.userRepository.CreateUser(&userData); err != nil {
		   c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create user: " + err.Error()})
		   return
	    }
 
	    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	}
 }
 
 
// Login handles user login.
func (ac *AuthController) Login()(gin.HandlerFunc) {

	return func(c *gin.Context) {
	// Implement user login logic here.
	// Check user credentials, generate tokens, etc.
  
	// Example response for successful login:
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
 
  
 // Logout handles user logout (if needed).
 
	// Implement user logout logic here (if needed).
  
	// Example response for successful logout:
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
 }

}
 