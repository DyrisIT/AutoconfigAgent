package handler

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/idna"
)

func validateEmail(email string) (string, error) {
	if email == "" {
		return "", fmt.Errorf("emailaddress parameter is missing")
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 || len(parts[0]) < 1 || len(parts[1]) < 1 {
		return "", fmt.Errorf("invalid email format")
	}

	// Convert domain part using IDNA punycode translation
	convertedDomain, err := idna.ToASCII(parts[1])
	if err != nil {
		return "", fmt.Errorf("invalid domain: %v", err)
	}

	if strings.Count(convertedDomain, ".") < 1 {
		return "", errors.New("domain must contain at least one period")
	}

	return parts[0] + "@" + convertedDomain, nil
}
