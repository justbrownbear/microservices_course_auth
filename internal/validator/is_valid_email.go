package validator

import "regexp"

// IsValidEmail checks if the provided email string is a valid email address
// according to a predefined regular expression pattern.
//
// Parameters:
//   - email: A string representing the email address to be validated.
//
// Returns:
//   - bool: True if the email address is valid, false otherwise.
func IsValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
