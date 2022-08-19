package pkg

import (
	"errors"
	"regexp"
	"strings"
)

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

// CheckValidEmails check valid for array of email address parameter
func CheckValidEmails(emails []string) error {
	if emails == nil || len(emails) == 0 {
		return errors.New("Email address is empty")
	}
	for _, email := range emails {
		if strings.TrimSpace(email) == "" || !IsEmailValid(email) {
			return errors.New("Invalid email address")
		}
	}
	return nil
}

// CheckValidEmail check valid for an email address parameter
func CheckValidEmail(email string) error {
	if strings.TrimSpace(email) == "" || !IsEmailValid(email) {
		return errors.New("Invalid email address")
	}
	return nil
}
