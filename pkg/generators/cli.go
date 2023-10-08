package generators

import (
	"github.com/jedi-knights/scaffit/pkg/fsys"
	"log"
)

// CliGenerator generates the cli structure
type CliGenerator struct {
	location string
}

// NewCliGenerator creates a new cli generator
func NewCliGenerator(location string) *CliGenerator {
	return &CliGenerator{
		location: location,
	}
}

// Location returns the location of the cli
func (g *CliGenerator) Location() string {
	return g.location
}

func (g *CliGenerator) generateFiles() error {
	log.Println("Todo: Generate cli files")
	return nil
}

// Generate generates the cli structure
func (g *CliGenerator) Generate() error {
	var (
		err error
	)

	log.Printf("Generating cli at %s\n", g.location)

	if !fsys.DirectoryExists(g.location) {
		log.Printf("Creating directory %s\n", g.location)
		if err = fsys.CreateDirectory(g.location); err != nil {
			return err
		}
	} else {
		log.Printf("Directory %s already exists\n", g.location)
	}

	if err = g.generateFiles(); err != nil {
		return err
	}

	return nil
}
