package utils

import (
	"fmt"
	"strings"
)

// ValidateModulePath validates the input string to ensure it is a valid module name.
//
// Module names should be in lowercase letters and avoid hyphens (-) or special characters.
// Use underscores or camelCase for multi-word module names. For example, mylibrary or
// myLibrary.
func ValidateModulePath(input string) error {
	if len(input) < 1 {
		return fmt.Errorf("Module name must be at least 1 character")
	}

	// Check to ensure the input is all lowercase
	if input != strings.ToLower(input) {
		return fmt.Errorf("Module name must be lowercase")
	}

	// Check to ensure the input does not contain hyphens
	if strings.Contains(input, "-") {
		return fmt.Errorf("Module name cannot contain hyphens")
	}

	// Check to ensure the input does not contain special characters
	if strings.ContainsAny(input, "!@#$%^&*()+=-`~[]{}\\|;:'\",.<>?") {
		return fmt.Errorf("Module name cannot contain special characters")
	}

	return nil
}
