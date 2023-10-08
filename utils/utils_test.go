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
		})

		Context("when the input is invalid", func() {
			It("returns an error", func() {
			})
		})
	})
})
