package pkg

import (
	"errors"
	"regexp"
	"strings"
)

// IsEmailValid function used for checking whether an email address is valid or not
// pass a string as parameter
// return a boolean
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

// CheckValidEmails used for checking an array of email address parameter are valid or not
// pass an email array as parameter
// return an error type
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

// CheckValidEmail used for checking whether an email address parameter is valid or not
// pass an email string as parameter
// return an error type
func CheckValidEmail(email string) error {
	if strings.TrimSpace(email) == "" || !IsEmailValid(email) {
		return errors.New("Invalid email address")
	}
	return nil
}
