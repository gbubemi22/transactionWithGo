package main

import (
    "fmt"
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/gbubemi22/transaction/src/model"
   // "github.com/joho/godotenv"
   // "os"
)


func init() {
    // Directly specify the database connection details
    dbHost := "localhost"
    dbPort := "5432"
    dbUser := "gbubemi"
    dbName := "gomyjob"
    dbPassword := "12345"
    dbSSLMode := "disable"
    dbSearchPath := "public"

    // Construct the database connection string
    connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s search_path=%s",
        dbHost, dbPort, dbUser, dbName, dbPassword, dbSSLMode, dbSearchPath)

    // Open a database connection
    db, err := gorm.Open("postgres", connectionString)
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }

    // AutoMigrate the User model
    db.AutoMigrate(&model.User{})
    db.AutoMigrate(&model.Wallet{})
    db.AutoMigrate(&model.Transaction{})
    fmt.Println("Migration complete")
}


// func init() {
//     // Load environment variables from .env
//     err := godotenv.Load()
//     if err != nil {
//         log.Fatal("Error loading .env file:", err)
//     }

//     // Construct the database connection string
//     dbHost := os.Getenv("DB_HOST")
//     dbPort := os.Getenv("DB_PORT")
//     dbUser := os.Getenv("DB_USER")
//     dbName := os.Getenv("DB_NAME")
//     dbPassword := os.Getenv("DB_PASSWORD")
//     dbSSLMode := os.Getenv("DB_SSLMODE")
//     dbSearchPath := os.Getenv("DB_SEARCH_PATH")

//     connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s search_path=%s",
//         dbHost, dbPort, dbUser, dbName, dbPassword, dbSSLMode, dbSearchPath)

//     // Open a database connection
//     db, err := gorm.Open("postgres", connectionString)
//     if err != nil {
//         log.Fatal("Error connecting to the database:", err)
//     }

//     // AutoMigrate the User model
//     db.AutoMigrate(&model.User{})
//     db.AutoMigrate(&model.Wallet{})
//     fmt.Println("Migration complete")
// }

func main() {
    // Your application logic can go here
}
