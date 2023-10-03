package db

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/gbubemi22/transaction/src/model"
    "github.com/joho/godotenv"
    "os"
)

func InitDB() (*gorm.DB, error) {
    // Load environment variables from .env
    err := godotenv.Load()
    if err != nil {
        return nil, err
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

    // Open a database connection
    db, err := gorm.Open("postgres", connectionString)
    if err != nil {
        return nil, err
    } 

    // AutoMigrate the User model
    if err := db.AutoMigrate(&model.User{}).Error; err != nil {
        return nil, err
    }

    return db, nil
}

if err := db.AutoMigrate(&model.Wallet{}).Error; err != nil {
    return nil, err
}

return db, nil









// package db

// import (
//     "fmt"
//     "github.com/jinzhu/gorm"
//     _ "github.com/jinzhu/gorm/dialects/postgres"
//     "github.com/gbubemi22/transaction/src/model"
//     "github.com/joho/godotenv"
//     "os"
// )

// func InitDB() (*gorm.DB, error) {
//     // Load environment variables from .env
//     err := godotenv.Load()
//     if err != nil {
//         return nil, err
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
//         return nil, err
//     } 
    

//     // AutoMigrate the User model
//     db.AutoMigrate(&model.User{})

   
  

//     return db, nil
// }
