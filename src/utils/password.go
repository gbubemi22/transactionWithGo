package utils



import (
	
	"golang.org/x/crypto/bcrypt"
 )


 // HashPassword hashes a password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
	    return "", err
	}
	return string(hashedPassword), nil
 }


 // VerifyPassword checks if the provided password matches the hashed password.


func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
	    return false, "Invalid credentials for password"
	}
	return true, ""
 }