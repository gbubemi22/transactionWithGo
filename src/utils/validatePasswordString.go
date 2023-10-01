package utils

import (
	"fmt"
	"regexp"
 )
 
 func ValidatePasswordString(password string) error {
	// Check for at least one digit
	hasDigit, _ := regexp.MatchString(`[0-9]`, password)
	if !hasDigit {
	    return fmt.Errorf("Password must contain at least one digit.")
	}
 
	// Check for at least one lowercase letter
	hasLowercase, _ := regexp.MatchString(`[a-z]`, password)
	if !hasLowercase {
	    return fmt.Errorf("Password must contain at least one lowercase letter.")
	}
 
	// Check for at least one uppercase letter
	hasUppercase, _ := regexp.MatchString(`[A-Z]`, password)
	if !hasUppercase {
	    return fmt.Errorf("Password must contain at least one uppercase letter.")
	}
 
	// Check for length between 8 and 20 characters
	if len(password) < 8 || len(password) > 20 {
	    return fmt.Errorf("Password must be between 8 and 20 characters long.")
	}
 
	return nil
 }
 