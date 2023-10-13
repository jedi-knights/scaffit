package pkg

import (
	"regexp"
	"strings"
)

func ContainsSpecialCharacters(part string) bool {
	// Regular expression pattern to check for special characters, allowing underscores
	// The pattern allows special characters !@#$%^&*()+=-`~[]{}\\|;:'\",<>\\(\\)?_.
	specialCharacters := "!@#$%^&*()+=`~[]{}\\\\|;:'\\\",<>\\\\(\\\\)?"

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
	pattern := "^[a-z]+(?:[a-zA-Z0-9-])*$"
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
