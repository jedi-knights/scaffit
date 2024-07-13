package pkg

import "testing"

func TestContainsSpecialCharacters(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "'my@project' contains a special character",
			input:    "my@project",
			expected: true,
		},
		{
			name:     "'myproject' does not contain special characters",
			input:    "myproject",
			expected: false,
		},
		{
			name:     "'my_project' does not contain special character",
			input:    "my_project",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			actual := ContainsSpecialCharacters(tt.input)

			// Assert
			if actual != tt.expected {
				t.Errorf("containsSpecialCharacters() = %v with '%s' want %v", actual, tt.input, tt.expected)
			}
		})
	}
}

func TestIsCamelCase(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "is camel case",
			input:    "myProject",
			expected: true,
		},
		{
			name:     "is not camel case",
			input:    "my_project",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			actual := IsCamelCase(tt.input)

			// Assert
			if actual != tt.expected {
				t.Errorf("isCamelCase() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestContainsWhitespace(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "contains whitespace",
			input:    "my project",
			expected: true,
		},
		{
			name:     "does not contain whitespace",
			input:    "myproject",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			actual := ContainsWhitespace(tt.input)

			// Assert
			if actual != tt.expected {
				t.Errorf("containsWhitespace() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
