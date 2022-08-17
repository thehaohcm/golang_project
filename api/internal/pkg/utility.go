package pkg

import (
	"errors"
	"golang_project/api/internal/models"
	"regexp"
	"strings"
)

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func GetDifference(relationship1 []models.Relationship, relationship2 []models.Relationship) []models.Relationship {
	var diff []models.Relationship

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range relationship1 {
			found := false
			for _, s2 := range relationship2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			relationship1, relationship2 = relationship2, relationship1
		}
	}

	return diff
}

func CheckValidEmails(emails []string) (bool, error) {
	if emails == nil {
		return false, errors.New("email address is emtpy")
	}
	for _, email := range emails {
		if strings.TrimSpace(email) == "" || !IsEmailValid(email) {
			return false, errors.New("invalid email address")
		}
	}
	return true, nil
}

func CheckValidEmail(email string) (bool, error) {
	if strings.TrimSpace(email) == "" || !IsEmailValid(email) {
		return false, errors.New("invalid email address")
	}
	return true, nil
}

func RemoveDuplicatedItems(items []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range items {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func RemoveItemInArray(l []string, item string) []string {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}
