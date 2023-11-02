package handler

import (
	"fmt"
	"regexp"
)

func validateEmail(email string) (string, error) {
	if email == "" {
		return "", fmt.Errorf("emailaddress parameter is missing")
	}

	// Regular expression to check if the email format is valid
	// Note: Email regex validation is complex and this simple regex does not cover all cases.
	re := regexp.MustCompile(`^.+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
	if !re.MatchString(email) {
		return "", fmt.Errorf("invalid email format")
	}

	return email, nil
}
