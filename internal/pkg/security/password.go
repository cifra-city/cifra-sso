package security

import "strings"

func HasRequiredChars(password string) bool {
	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ", char):
			hasUpper = true
		case strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", char):
			hasLower = true
		case strings.ContainsRune("0123456789", char):
			hasNumber = true
		case strings.ContainsRune("!@#$%^&*()-_+=<>?", char):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}
