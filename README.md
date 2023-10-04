Transaction Service

This is a Go application that provides a transaction service for transferring funds between user wallets. It uses the Gin framework for handling HTTP requests and the GORM library for database operations.
Prerequisites

Before running the application, make sure you have the following prerequisites installed:

    Go (https://golang.org/)
    PostgreSQL (https://www.postgresql.org/)

Installation

    Clone the repository to your local machine:

    bash

git clone https://github.com/gbubemi22/transactionWithGo.git

Change the directory to the project folder:



Install the required dependencies using Go Modules:

bash

go mod tidy

Configure the database connection by updating .env 
go file with your PostgreSQL database credentials.

Create the necessary database tables by running the database migration:

bash

cd src/migrations
    go run migration.go

Usage

    Start the application:

    bash

go run main.go

The application will run on http://localhost:8080 by default. You can customize the port and other settings in the config/config.go file.

To transfer funds between user wallets, make a POST request to /api/v1/transactions/transfer with the following JSON payload:

json

    {
      "fromUsername": "sender_username",
      "toUsername": "receiver_username",
      "amount": 100.0,
      "summary": "Payment for order #123"
    }

    Replace sender_username, receiver_username, and other fields as needed.

    The response will indicate the success or failure of the transaction.

API Endpoints

    POST /api/v1/transactions/transfer: Transfer funds between user wallets.

     POST /api/v1/auth/register:  Create a new account


      POST /api/v1/wallet/:userId:: To create wallet with a default blance of 0.00


      

Error Handling

    Invalid JSON data will result in a 400 Bad Request response.
    Insufficient balance will result in a 400 Bad Request response.
    User not found will result in a 404 Not Found response.
    Database errors will result in a 500 Internal Server Error response.

