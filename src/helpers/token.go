package helpers

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

// Define the secret key used to sign the JWT tokens.
var jwtSecret = []byte("your_secret_key_here")

// GenerateJWT generates a new JWT token for the given user ID.
func GenerateJWT(userID uint) (string, error) {
    // Define the token claims
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (1 day)
        "iat":     time.Now().Unix(),                     // Token issuance time
    }

    // Create the token with the claims and sign it using the secret key
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// VerifyJWT verifies and parses a JWT token, returning the user ID if valid.
func VerifyJWT(tokenString string) (uint, error) {
    // Parse the token with the secret key
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil {
        return 0, err
    }

    // Check if the token is valid
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userID := uint(claims["user_id"].(float64))
        return userID, nil
    }

    return 0, jwt.ErrSignatureInvalid
}
