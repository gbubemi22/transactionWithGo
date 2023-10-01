package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/gbubemi22/transaction/src/model"
    "github.com/gbubemi22/transaction/src/utils"
    "github.com/gbubemi22/transaction/src/helpers"
    "github.com/jinzhu/gorm"
    "net/http"
    "fmt"
    
)

type AuthController struct {
    db *gorm.DB
}

type ErrorResponse struct {
    Error string `json:"error"`
}

func NewAuthController(db *gorm.DB) *AuthController {
    return &AuthController{db}
}

func (ac *AuthController) Signup() gin.HandlerFunc {
    return func(c *gin.Context) {
        var userData model.User
        if err := c.ShouldBindJSON(&userData); err != nil {
            c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON data: " + err.Error()})
            return
        }

        

        // Check if the email already exists in the database.
        var existingUser model.User
        if err := ac.db.Where("email = ?", userData.Email).First(&existingUser).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        fmt.Println(existingUser)
        if existingUser.ID != 0 {
            c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
            return
        }

        // Check if the phone number already exists in the database.
        if err := ac.db.Where("phone = ?", userData.Phone).First(&existingUser).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        if existingUser.ID != 0 {
            c.JSON(http.StatusConflict, gin.H{"error": "Phone number already exists"})
            return
        }
	  
        // Validate the user's password using the correct function name.
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

        if err := ac.db.Create(&userData).Error; err != nil {
            c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create user: " + err.Error()})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
    }
}



func (ac *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
	    // Retrieve the user by email from the database
	    var userData model.User
 
	    // Parse user login credentials from the request body
	    if err := c.ShouldBindJSON(&userData); err != nil {
		   c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON data: " + err.Error()})
		   return
	    }
 
	    // Get the user by email
	    var user model.User
	    	    if err := ac.db.Where("email = ?", userData.Email).First(&user).Error; err != nil {
	    		   c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid credentials for email"})
	    		   return
	    	    }
 
	//     Verify the provided password with the stored hashed password
	    	    passwordsMatch, errMsg := utils.VerifyPassword(user.Password, userData.Password)
	    if !passwordsMatch {
		   c.JSON(http.StatusUnauthorized, ErrorResponse{Error: errMsg})
		   return
	    }
 
	    // In your AuthController's Login function
	    token, err := helpers.GenerateJWT(user.ID)
	    if err != nil {
		   c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to generate token"})
		   return
	    }
 
	    c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"first_name":  user.FirstName,
			"last_name":   user.LastName,
			"email":       user.Email,
			"phone":       user.Phone,
	})
	}
 }
 
// func (ac *AuthController) Login() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 	    var loginData struct {
// 		   Email    string `json:"email"`
// 		   Password string `json:"password"`
// 	    }
 
// 	    // Parse user login credentials from the request body
// 	    if err := c.ShouldBindJSON(&loginData); err != nil {
// 		   c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON data: " + err.Error()})
// 		   return
// 	    }
 
// 	    // Find the user by email in the database
// 	    var user model.User
// 	    if err := ac.db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
// 		   c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid credentials for email"})
// 		   return
// 	    }


// 	    fmt.Printf("Entered Email: %#v\n", loginData.Email)
// fmt.Printf("Stored Email: %#v\n", user.Email)
// fmt.Printf("Entered Password: %#v\n", loginData.Password)




// 	//     if loginData.Password != user.Password {
// 	// 	fmt.Println("Provided Password:", loginData.Password)
// 	// 	fmt.Println("Stored Hashed Password:", user.Password)
// 	// 	c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid credentials for password check"})
// 	// 	return
// 	//  }
 
// 	    // Verify the provided password with the hashed password from the database
// 	    passwordsMatch, errMsg := utils.VerifyPassword(user.Password, loginData.Password)
// 	    if !passwordsMatch {
// 		   c.JSON(http.StatusUnauthorized, ErrorResponse{Error: errMsg})
// 		   return
// 	    }
	    
  	    
 
	   
//            token, err := helpers.GenerateJWT(user.ID)
//              if err != nil {
//             c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to generate token"})
//             return
//           }
 
// 	    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
// 	}
//  }
 
