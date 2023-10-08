package utils_test

import (
	"github.com/jedi-knights/scaffit/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	Describe("ValidateModulePath", func() {
		Context("when the input is valid", func() {
			It("returns nil when given 'github.com/jdoe/myproject", func() {
				// Arrange
				modulePath := "github.com/jdoe/myproject"

				// Act
				err := utils.ValidateModulePath(modulePath)

				// Assert
				Expect(err).To(BeNil())
			})

			It("returns nil when given 'github.com/jdoe/my_project", func() {
				// Arrange
				modulePath := "github.com/jdoe/my_project"

				// Act
				err := utils.ValidateModulePath(modulePath)

				// Assert
				Expect(err).To(BeNil())
			})

			It("returns nil when given 'github.com/jdoe/myLibrary", func() {
				// Arrange
				modulePath := "github.com/jdoe/myLibrary"

				// Act
				err := utils.ValidateModulePath(modulePath)

				// Assert
				Expect(err).To(BeNil())
			})
		})

		Context("when the input is invalid", func() {
			It("returns an error when given a string with spaces", func() {
				// Arrange
				modulePath := "github.com/jdoe/my project"

				// Act
				err := utils.ValidateModulePath(modulePath)

				// Assert
				Expect(err).ToNot(BeNil())
			})

			It("returns an error when given a string with a hyphen", func() {
				// Arrange
				modulePath := "github.com/jdoe/my-project"

				// Act
				err := utils.ValidateModulePath(modulePath)

				// Assert
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("ContainsSpecialCharacters", func() {
		It("returns true when given a string with a special character", func() {
			// Arrange
			part := "my@project"

			// Act
			result := utils.ContainsSpecialCharacters(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns false when given a string without a special character", func() {
			// Arrange
			part := "myproject"

			// Act
			result := utils.ContainsSpecialCharacters(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given my_project", func() {
			// Arrange
			part := "my_project"

			// Act
			result := utils.ContainsSpecialCharacters(part)

			// Assert
			Expect(result).To(BeFalse())
		})
	})

	Describe("IsCamelCase", func() {
		It("returns true when given a string in camel case", func() {
			// Arrange
			part := "myProject"

			// Act
			result := utils.IsCamelCase(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns false when given a string not in camel case", func() {
			// Arrange
			part := "my_project"

			// Act
			result := utils.IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given a string with a hyphen", func() {
			// Arrange
			part := "my-project"

			// Act
			result := utils.IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given a string with a special character", func() {
			// Arrange
			part := "my_project"

			// Act
			result := utils.IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given a string with a space", func() {
			// Arrange
			part := "my project"

			// Act
			result := utils.IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})

		It("returns false when given a string with an uppercase letter", func() {
			// Arrange
			part := "myProject"

			// Act
			result := utils.IsCamelCase(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns false when given a string with a number", func() {
			// Arrange
			part := "myProject1"

			// Act
			result := utils.IsCamelCase(part)

			// Assert
			Expect(result).To(BeFalse())
		})
	})

	Describe("ContainsWhitespace", func() {
		It("returns true when given a string with a space", func() {
			// Arrange
			part := "my project"

			// Act
			result := utils.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a tab", func() {
			// Arrange
			part := "my	project"

			// Act
			result := utils.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a newline", func() {
			// Arrange
			part := "my\nproject"

			// Act
			result := utils.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a carriage return", func() {
			// Arrange
			part := "my\rproject"

			// Act
			result := utils.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a form feed", func() {
			// Arrange
			part := "my\fproject"

			// Act
			result := utils.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns true when given a string with a vertical tab", func() {
			// Arrange
			part := "my\vproject"

			// Act
			result := utils.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeTrue())
		})

		It("returns false when given a string without a whitespace character", func() {
			// Arrange
			part := "myproject"

			// Act
			result := utils.ContainsWhitespace(part)

			// Assert
			Expect(result).To(BeFalse())
		})
	})
})
