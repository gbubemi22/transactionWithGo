package utils


import (
	"fmt"
	"regexp"
 )
 
 // validatePasswordString validates a password string.
 func ValidatePasswordString(password string) error {
	
	pattern := `^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,20}$`
 
	// Compile the regular expression pattern.
	regex, err := regexp.Compile(pattern)
	if err != nil {
	    return err
	}
 
		if !regex.MatchString(password) {
	    return fmt.Errorf("Password must contain a capital letter, number, special character, and be between 8 and 20 characters long.")
	}
 
	return nil
 }
 

 