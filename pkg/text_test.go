package pkg_test

import (
	. "github.com/jedi-knights/scaffit/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Text", func() {
	Describe("ContainsSpecialCharacters", func() {
		It("returns true when given a string with a special character", func() {
			// Arrange
			part := "my@project"

			// Act
			result := ContainsSpecialCharacters(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns false when given a string without a special character", func() {
			// Arrange
			part := "myproject"

			// Act
			result := ContainsSpecialCharacters(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given my_project", func() {
			// Arrange
			part := "my_project"

			// Act
			result := ContainsSpecialCharacters(part)

			// Assert
			Expect(result).To(BeFalse())
		})
	})

	Describe("IsCamelCase", func() {
		It("returns true when given a string in camel case", func() {
			// Arrange
			part := "myProject"

			// Act
			result := IsCamelCase(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns false when given a string not in camel case", func() {
			// Arrange
			part := "my_project"

			// Act
			result := IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given a string with a hyphen", func() {
			// Arrange
			part := "my-project"

			// Act
			result := IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given a string with a special character", func() {
			// Arrange
			part := "my_project"

			// Act
			result := IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given a string with a space", func() {
			// Arrange
			part := "my project"

			// Act
			result := IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given a string with an uppercase letter", func() {
			// Arrange
			part := "myProject"

			// Act
			result := IsCamelCase(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a camel cased string with a number", func() {
			// Arrange
			part := "myProject1"

			// Act
			result := IsCamelCase(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given 'project1'", func() {
			// Arrange
			part := "project1"

			// Act
			result := IsCamelCase(part)

			// Assert
			Expect(result).To(BeTrue())
		})
	})

	Describe("ContainsWhitespace", func() {
		It("returns true when given a string with a space", func() {
			// Arrange
			part := "my project"

			// Act
			result := ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a tab", func() {
			// Arrange
			part := "my	project"

			// Act
			result := ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a newline", func() {
			// Arrange
			part := "my\nproject"

			// Act
			result := ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a carriage return", func() {
			// Arrange
			part := "my\rproject"

			// Act
			result := ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a form feed", func() {
			// Arrange
			part := "my\fproject"

			// Act
			result := ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a vertical tab", func() {
			// Arrange
			part := "my\vproject"

			// Act
			result := ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns false when given a string without a whitespace character", func() {
			// Arrange
			part := "myproject"

			// Act
			result := ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeFalse())
		})
	})
})
