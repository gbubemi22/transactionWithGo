package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "log"
    "fmt"
    "os"
    "github.com/gbubemi22/transaction/src/model"
    "github.com/gbubemi22/transaction/src/routes"
    "github.com/gbubemi22/transaction/src/controller"
    //"github.com/gbubemi22/transaction/src/utils"
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


    // Construct the database connection string
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbName := os.Getenv("DB_NAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbSSLMode := os.Getenv("DB_SSLMODE")

    connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
        dbHost, dbPort, dbUser, dbName, dbPassword, dbSSLMode)

    // Initialize the database connection
    db, err := gorm.Open("postgres", connectionString)
    if err != nil {
        panic(err)
    } else {
        fmt.Println("Connected to the database successfully")
    }

    defer db.Close()

    PerformMigrations(db)

   
 

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    router := gin.New()
    router.Use(gin.Logger())


    authController := controller.NewAuthController(db) 
    userController := controller.NewUserController(db) 

    routes.AuthRoutes(router,authController )
    routes.UserRoutes(router, userController)

    

    router.GET("/", func(c *gin.Context) {
        fmt.Println("Welcome to my application")
        c.JSON(200, gin.H{"message": "Welcome to my application"})
    })

    router.Run(":" + port)
}



// // Set up authentication-related routes.
    // authController := NewAuthController(db)
    // AuthRoutes(router, authController)

    // // Set up user-related routes.
    // userController := NewUserController(db)
    // UserRoutes(router, userController)









  