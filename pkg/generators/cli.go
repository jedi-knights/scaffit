package generators

import (
	"github.com/jedi-knights/scaffit/pkg"
	"log"
)

// CliGenerator generates the cli structure
type CliGenerator struct {
	fsys     pkg.FileSystem
	location string
}

// NewCliGenerator creates a new cli generator
func NewCliGenerator(fsys pkg.FileSystem, location string) *CliGenerator {
	return &CliGenerator{
		fsys:     fsys,
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

	if !g.fsys.DirectoryExists(g.location) {
		log.Printf("Creating directory %s\n", g.location)
		if err = g.fsys.CreateDirectory(g.location, false); err != nil {
			return err
		}
	} else {
		log.Printf("Directory %s exists\n", g.location)
	}

	if err = g.generateFiles(); err != nil {
		return err
	}

	return nil
}
