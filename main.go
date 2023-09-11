package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/jinzhu/gorm"
    "github.com/gbubemi22/transaction/src/model"
    DB "github.com/gbubemi22/transaction/src/DB"
    routes "github.com/gbubemi22/transaction/src/routers"
    "github.com/gbubemi22/transaction/src/controller"
    "github.com/gbubemi22/transaction/src/repository"
    "log"
    "fmt"
    "os"
)

func PerformMigrations(db *gorm.DB) {
    // Apply all available migrations.
    db.AutoMigrate(&model.User{})
}

func main() {
    // Load environment variables from .env file
    loadErr := godotenv.Load(".env")
    if loadErr != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize the database connection
    db, err := DB.InitDB()
    if err != nil {
        panic(err)
    } else {
        fmt.Println("connected to the database successfully")
    }
    
    defer db.Close()

    PerformMigrations(db)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    router := gin.New()
    router.Use(gin.Logger())

    // Create the UserRepository with the database connection
     //userRepository := repository.NewUserRepository(db, "", "", "","", "")


   userRepository := repository.NewUserRepository(db)


    authController := controller.NewAuthController(userRepository)

    userController := controller.NewUserController(userRepository)

	// Set up user-related routes.
	routes.UserRoutes(router, userController)

    // Set up authentication-related routes.
    routes.AuthRoutes(router, authController)

    router.GET("/", func(c *gin.Context) {
        fmt.Println("Welcome to my application")
        c.JSON(200, gin.H{"message": "Welcome to my application"})
    })

    router.Run(":" + port)
}
