package generators

import (
	"github.com/jedi-knights/scaffit/pkg/fsys"
	"github.com/jedi-knights/scaffit/utils"
	"github.com/manifoldco/promptui"
	"log"
)

// ModuleGenerator generates the module structure
type ModuleGenerator struct {
	location string
}

// NewModuleGenerator creates a new module generator
func NewModuleGenerator(location string) *ModuleGenerator {
	return &ModuleGenerator{
		location: location,
	}
}

// Location returns the location of the module
func (g *ModuleGenerator) Location() string {
	return g.location
}

func (g *ModuleGenerator) generateFiles() error {
	var (
		err        error
		modulePath string
	)

	prompt := promptui.Prompt{
		Label:    "Module path",
		Validate: utils.ValidateModulePath,
	}

	if modulePath, err = prompt.Run(); err != nil {
		return err
	}

	log.Println("Todo: Generate module files")
	return nil
}

// Generate generates the module structure
func (g *ModuleGenerator) Generate() error {
	log.Printf("Generating module at %s\n", g.location)

	if !fsys.DirectoryExists(g.location) {
		log.Printf("Creating directory %s\n", g.location)
		if err := fsys.CreateDirectory(g.location); err != nil {
			return err
		}
	} else {
		log.Printf("Directory %s already exists\n", g.location)
	}

	return nil
}
