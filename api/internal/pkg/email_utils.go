package pkg

import (
	"errors"
	"regexp"
	"strings"
)

// CheckValidEmails used for checking an array of email address parameter are valid or not
// pass an email array as parameter
// return an error type
func CheckValidEmails(emails []string) error {
	if emails == nil || len(emails) == 0 {
		return errors.New("email address is empty")
	}
	for _, email := range emails {
		if CheckValidEmail(email) != nil {
			return errors.New("invalid email address")
		}
	}

	return nil
}

// CheckValidEmail used for checking whether an email address parameter is valid or not
// pass an email string as parameter
// return an error type
func CheckValidEmail(email string) error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if strings.TrimSpace(email) == "" || !emailRegex.MatchString(email) {
		return errors.New("invalid email address")
	}

	return nil
}
