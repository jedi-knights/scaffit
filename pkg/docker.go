// This module is responsible for generating Docker files.

package pkg

import "fmt"

var NotImplementedError = fmt.Errorf("not implemented")

type FileGenerator interface {
	Generate() error
}

// DockerFileGenerator creates a new instance of the Dockerfile
type DockerFileGenerator struct {
	FileGenerator
}

// DockerIgnoreFileGenerator creates a new instance of the ignore list
type DockerIgnoreFileGenerator struct {
	FileGenerator
}

func (g DockerFileGenerator) Generate() error {
	return NotImplementedError
}

func (g DockerIgnoreFileGenerator) Generate() error {
	return NotImplementedError
}
