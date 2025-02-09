package utils

import "log"

func ValidatePassword(password string) error {
	if len(password) < 5 {
		log.Printf("Password must be at least 5 characters long: %v", password)
		return nil
	}

	return nil
}
