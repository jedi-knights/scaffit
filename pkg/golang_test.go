package pkg_test

//var _ = Describe("Golang", func() {
//	Describe("ValidateModulePath", func() {
//		Context("when the userInput is valid", func() {
//			It("returns nil when given 'github.com/jdoe/myproject", func() {
//				// Arrange
//				modulePath := "github.com/jdoe/myproject"
//
//				// Act
//				err := pkg.ValidateModulePath(modulePath)
//
//				// Assert
//				Expect(err).To(BeNil())
//			})
//
//			It("returns nil when given 'github.com/jedi-knights/rankings", func() {
//				// Arrange
//				modulePath := "github.com/jedi-knights/rankings"
//
//				// Act
//				err := pkg.ValidateModulePath(modulePath)
//
//				// Assert
//				Expect(err).To(BeNil())
//			})
//
//			It("returns nil when given 'github.com/jdoe/my_project", func() {
//				// Arrange
//				modulePath := "github.com/jdoe/my_project"
//
//				// Act
//				err := pkg.ValidateModulePath(modulePath)
//
//				// Assert
//				Expect(err).To(BeNil())
//			})
//
//			It("returns nil when given 'github.com/jdoe/myLibrary", func() {
//				// Arrange
//				modulePath := "github.com/jdoe/myLibrary"
//
//				// Act
//				err := pkg.ValidateModulePath(modulePath)
//
//				// Assert
//				Expect(err).To(BeNil())
//			})
//		})
//
//		Context("when the userInput is invalid", func() {
//			It("returns an error when given a string with spaces", func() {
//				// Arrange
//				modulePath := "github.com/jdoe/my project"
//
//				// Act
//				err := pkg.ValidateModulePath(modulePath)
//
//				// Assert
//				Expect(err).ToNot(BeNil())
//			})
//
//			It("returns an error when given a string with a hyphen", func() {
//				// Arrange
//				modulePath := "github.com/jdoe/my-project"
//
//				// Act
//				err := pkg.ValidateModulePath(modulePath)
//
//				// Assert
//				Expect(err).ToNot(BeNil())
//			})
//		})
//	})
//})
