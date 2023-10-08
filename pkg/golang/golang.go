package golang

import (
	"fmt"
	"github.com/jedi-knights/scaffit/pkg/text"
	"os"
	"os/exec"
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
		if text.ContainsSpecialCharacters(part) {
			return fmt.Errorf("Module path part contains special characters: %s", part)
		}

		if text.ContainsWhitespace(part) {
			return fmt.Errorf("Module path part contains whitespace characters: %s", part)
		}

		if !strings.Contains(part, ".") && !strings.Contains(part, "_") && !text.IsCamelCase(part) {
			return fmt.Errorf("Module path part is not in camel case: %s", part)
		}
	}

	return nil
}

func InitializeGoModule(location, modulePath string) error {
	// Command and arguments to run "go mod init" with the specified module path
	cmd := exec.Command("go", "mod", "init", modulePath)

	// Set the working directory for the command (optional)
	cmd.Dir = location

	// Redirect command output to the standard streams
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
