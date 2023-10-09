package generators

import (
	"github.com/jedi-knights/scaffit/pkg"
	"log"
)

// ApiGenerator generates the api structure
type ApiGenerator struct {
	fsys     pkg.FileSystem
	location string
}

func NewApiGenerator(fsys pkg.FileSystem, location string) *ApiGenerator {
	return &ApiGenerator{
		fsys:     fsys,
		location: location,
	}
}

// Location returns the location of the api
func (g *ApiGenerator) Location() string {
	return g.location
}

func (g *ApiGenerator) generateFiles() error {
	log.Println("Todo: Generate api files")
	return nil
}

// Generate generates the api structure
func (g *ApiGenerator) Generate() error {
	var (
		err error
	)

	log.Printf("Generating api at %s\n", g.location)

	if !g.fsys.DirectoryExists(g.location) {
		if err = g.fsys.CreateDirectory(g.location); err != nil {
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
