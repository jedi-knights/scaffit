package generators

import (
	"github.com/jedi-knights/scaffit/pkg/fsys"
	"log"
)

// ApiGenerator generates the api structure
type ApiGenerator struct {
	location string
}

func NewApiGenerator(location string) *ApiGenerator {
	return &ApiGenerator{
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

	if !fsys.DirectoryExists(g.location) {
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
