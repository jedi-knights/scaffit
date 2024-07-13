package pkg

import (
	"github.com/jedi-knights/scaffit/pkg/input"
	"strings"
	"testing"
)

// NewApplication test
func TestNewApplication(t *testing.T) {
	// Arrange
	in := input.MockedUserInput{}

	// Act
	app := NewApplication(in)

	// Assert
	if app == nil {
		t.Fatalf("NewApplication should return a non-nil Application instance")
	}

	for _, f := range Frameworks {
		if _, ok := app.useFlags[strings.ToLower(f)]; !ok {
			t.Errorf("useFlags map should contain key %s", f)
		}
	}

	if len(app.useFlags) != len(Frameworks) {
		t.Errorf("useFlags map should have %d keys", len(Frameworks))
	}
}

type modulePathError struct {
	part string
	msg  string
}

func (m modulePathError) Error() string {
	return m.msg
}

func TestValidateModulePath(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    string
		expected error
	}{
		{
			name:     "valid module path",
			input:    "github.com/jdoe/myproject",
			expected: nil,
		},
		{
			name:     "valid module path",
			input:    "github.com/jedi-knights/rankings",
			expected: nil,
		},
		{
			name:     "valid module path",
			input:    "github.com/jdoe/my_project",
			expected: nil,
		},
		{
			name:     "valid module path",
			input:    "github.com/jdoe/myLibrary",
			expected: nil,
		},
		{
			name:  "invalid module path",
			input: "github.com/jdoe/my project",
			expected: &modulePathError{
				part: "my project",
				msg:  "module path part contains whitespace characters: my project",
			},
		},
		{
			name:  "invalid module path",
			input: "github.com/jdoe/my@project",
			expected: &modulePathError{
				part: "my-project",
				msg:  "module path part contains special characters: my@project",
			},
		},
		{
			name:  "invalid module path",
			input: "github.com/jdoe/123NumbersFirst",
			expected: &modulePathError{
				part: "123NumbersFirst",
				msg:  "module path part is not in camel case: '123NumbersFirst' else 'github.com/jdoe/123NumbersFirst'",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			err := ValidateModulePath(tt.input)

			// Assert
			if err != nil {
				if tt.expected == nil {
					t.Fatalf("expected nil error, got: %v", err)
				}

				if err.Error() != tt.expected.Error() {
					t.Fatalf("expected error '%v', got: '%v'", tt.expected, err)
				}
			} else if tt.expected != nil {
				t.Fatalf("expected error '%v', got nil for path '%s'", tt.expected, tt.input)
			}
		})
	}
}

func TestGetModulePath(t *testing.T) {
	// Arrange
	in := input.MockedUserInput{
		PromptResponses: map[string]string{
			"Module path": "github.com/jdoe/myproject",
		},
	}
	app := NewApplication(in)

	// Act
	modulePath, err := app.GetModulePath()

	// Assert
	if err != nil {
		t.Fatalf("GetModulePath() failed: %v", err)
	}

	if modulePath == "" {
		t.Fatalf("GetModulePath() should return a non-empty string")
	}
}
