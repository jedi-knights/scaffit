package pkg

import (
	"fmt"
	"github.com/jedi-knights/scaffit/pkg/input"
	"strings"
)

var Frameworks = []string{"Echo", "Gin", "Fiber", "Mux", "None"}

type Applicationer interface {
	GetModulePath() (string, error)
	GetFramework() (string, error)
}

type Application struct {
	userInput input.UserInput
	useFlags  map[string]bool
}

func NewApplication(in input.UserInput) *Application {
	app := Application{
		userInput: in,
		useFlags:  make(map[string]bool),
	}

	app.InitializeUseFlags()

	return &app
}

func (a Application) InitializeUseFlags() {
	for _, f := range Frameworks {
		a.useFlags[strings.ToLower(f)] = false
	}
}

// ValidateModulePath validates the userInput string to ensure it is a valid module name.
//
// Module names should be in lowercase letters and avoid hyphens (-) or special characters.
// Use underscores or camelCase for multi-word module names. For example, mylibrary or
// myLibrary.
func ValidateModulePath(input string) error {
	// Split the userInput string on '/'
	parts := strings.Split(input, "/")

	// Check each part individually
	for _, part := range parts {
		if ContainsSpecialCharacters(part) {
			return fmt.Errorf("module path part contains special characters: %s", part)
		}

		if ContainsWhitespace(part) {
			return fmt.Errorf("module path part contains whitespace characters: %s", part)
		}

		if !strings.Contains(part, ".") && !strings.Contains(part, "_") && !IsCamelCase(part) {
			return fmt.Errorf("module path part is not in camel case: '%s' else '%s'", part, input)
		}
	}

	return nil
}

func (a *Application) GetModulePath() (string, error) {
	modulePath, err := a.userInput.PromptForString("Module path", ValidateModulePath)
	if err != nil {
		return "", err
	}

	return modulePath, nil
}

func (a *Application) GetFramework() (string, error) {
	label := "Which framework would you like to use?"
	frameworks := []string{"Echo", "Gin", "Fiber", "Mux", "None"}

	_, framework, err := a.userInput.SelectFromItems(label, frameworks)
	if err != nil {
		return "", err
	}

	for _, f := range frameworks {
		a.useFlags[strings.ToLower(f)] = false
	}

	a.useFlags[strings.ToLower(framework)] = true

	return framework, nil
}
