package golang

import (
	"fmt"
	"github.com/jedi-knights/scaffit/pkg"
	"github.com/jedi-knights/scaffit/pkg/text"
	"log"
	"path/filepath"
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
			return fmt.Errorf("module path part contains special characters: %s", part)
		}

		if text.ContainsWhitespace(part) {
			return fmt.Errorf("module path part contains whitespace characters: %s", part)
		}

		if !strings.Contains(part, ".") && !strings.Contains(part, "_") && !text.IsCamelCase(part) {
			return fmt.Errorf("module path part is not in camel case: %s", part)
		}
	}

	return nil
}

func InitializeGoModule(location, modulePath string) error {
	var (
		err error
	)

	if err = pkg.RunCommand(location, "go mod init "+modulePath, false); err != nil {
		return err
	}
	if err = InstallDependencies(location); err != nil {
		return err
	}

	// Download the .gitignore file
	uri := "https://raw.githubusercontent.com/github/gitignore/master/Go.gitignore"
	gitIgnorePath := filepath.Join(location, ".gitignore")
	if err = pkg.DownloadFile(uri, gitIgnorePath); err != nil {
		return err
	}

	return nil
}

func InstallDependency(location string, dependency string) error {
	if err := pkg.RunCommand(location, "go get "+dependency, false); err != nil {
		return err
	}

	return nil
}

func InstallDependencies(location string) error {
	dependencies := []string{
		"github.com/onsi/ginkgo/v2",
		"github.com/onsi/gomega",
	}

	for _, dependency := range dependencies {
		if err := InstallDependency(location, dependency); err != nil {
			return err
		}
	}

	return nil
}

func InitializeCobra(location string, useViper bool) error {
	var (
		err     error
		command string
	)

	// Install Cobra
	if err = InstallDependency(location, "github.com/spf13/cobra"); err != nil {
		return err
	}

	if useViper {
		log.Print("Using Viper\n")
		// Install Viper
		if err = InstallDependency(location, "github.com/spf13/viper"); err != nil {
			return err
		}

		command = "cobra-cli init --viper"
	} else {
		log.Print("Not using Viper\n")
		command = "cobra-cli init"
	}

	if err = pkg.RunCommand(location, command, false); err != nil {
		return err
	}

	return nil
}
