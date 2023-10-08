package text_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jedi-knights/scaffit/pkg/text"
)

var _ = Describe("Text", func() {
	Describe("ContainsSpecialCharacters", func() {
		It("returns true when given a string with a special character", func() {
			// Arrange
			part := "my@project"

			// Act
			result := text.ContainsSpecialCharacters(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns false when given a string without a special character", func() {
			// Arrange
			part := "myproject"

			// Act
			result := text.ContainsSpecialCharacters(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given my_project", func() {
			// Arrange
			part := "my_project"

			// Act
			result := text.ContainsSpecialCharacters(part)

			// Assert
			Expect(result).To(BeFalse())
		})
	})

	Describe("IsCamelCase", func() {
		It("returns true when given a string in camel case", func() {
			// Arrange
			part := "myProject"

			// Act
			result := text.IsCamelCase(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns false when given a string not in camel case", func() {
			// Arrange
			part := "my_project"

			// Act
			result := text.IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given a string with a hyphen", func() {
			// Arrange
			part := "my-project"

			// Act
			result := text.IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given a string with a special character", func() {
			// Arrange
			part := "my_project"

			// Act
			result := text.IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given a string with a space", func() {
			// Arrange
			part := "my project"

			// Act
			result := text.IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given a string with an uppercase letter", func() {
			// Arrange
			part := "myProject"

			// Act
			result := text.IsCamelCase(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns false when given a string with a number", func() {
			// Arrange
			part := "myProject1"

			// Act
			result := text.IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})
	})

	Describe("ContainsWhitespace", func() {
		It("returns true when given a string with a space", func() {
			// Arrange
			part := "my project"

			// Act
			result := text.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a tab", func() {
			// Arrange
			part := "my	project"

			// Act
			result := text.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a newline", func() {
			// Arrange
			part := "my\nproject"

			// Act
			result := text.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a carriage return", func() {
			// Arrange
			part := "my\rproject"

			// Act
			result := text.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a form feed", func() {
			// Arrange
			part := "my\fproject"

			// Act
			result := text.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a vertical tab", func() {
			// Arrange
			part := "my\vproject"

			// Act
			result := text.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns false when given a string without a whitespace character", func() {
			// Arrange
			part := "myproject"

			// Act
			result := text.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeFalse())
		})
	})
})
