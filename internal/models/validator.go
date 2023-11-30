package models

import (
	"regexp"
)

func Valid(o Order) bool {
	d := o.Delivery

	if len(d.Phone) != 12 {
		return false
	}
	if !isValidEmail(d.Email) {
		return false
	}
	return true
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(email) {
		return true
	}
	return false
}
