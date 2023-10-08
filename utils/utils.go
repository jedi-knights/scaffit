package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateModulePath validates the input string to ensure it is a valid module name.
//
// Module names should be in lowercase letters and avoid hyphens (-) or special characters.
// Use underscores or camelCase for multi-word module names. For example, mylibrary or
// myLibrary.
func ValidateModulePath(input string) error {
	// Split the input string on '/'
	parts := strings.Split(input, "/")

	// Check each part individually
	for _, part := range parts {
		if ContainsSpecialCharacters(part) {
			return fmt.Errorf("Module path part contains special characters: %s", part)
		}

		if ContainsWhitespace(part) {
			return fmt.Errorf("Module path part contains whitespace characters: %s", part)
		}

		if !strings.Contains(part, ".") && !strings.Contains(part, "_") && !IsCamelCase(part) {
			return fmt.Errorf("Module path part is not in camel case: %s", part)
		}
	}

	return nil
}

func ContainsSpecialCharacters(part string) bool {
	// Regular expression pattern to check for special characters, allowing underscores
	// The pattern allows special characters !@#$%^&*()+=-`~[]{}\\|;:'\",<>\\(\\)?_.
	specialCharacters := "!@#$%^&*()+=-`~[]{}\\\\|;:'\\\",<>\\\\(\\\\)?"

	for _, char := range specialCharacters {
		if strings.Contains(part, string(char)) {
			return true
		}
	}

	return false
}

func IsCamelCase(part string) bool {
	// Regular expression pattern to check if a string is in camel case
	// The pattern checks for lowercase letters followed by uppercase letters.
	pattern := "^[a-z]+(?:[A-Z][a-z]*)*$"
	return regexp.MustCompile(pattern).MatchString(part)
}

func ContainsWhitespace(part string) bool {
	// Regular expression pattern to check for whitespace characters
	// The pattern checks for any whitespace character (space, tab, newline, etc.).
	whiteSpacePattern := `[ \t\v\n\r\f]`

	// Compile the regular expression
	re := regexp.MustCompile(whiteSpacePattern)

	// Use the regular expression to find whitespace characters in the input string
	return re.MatchString(part)
}
