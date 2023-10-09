package generators

import (
	"github.com/jedi-knights/scaffit/pkg"
	"github.com/jedi-knights/scaffit/pkg/golang"
	"github.com/jedi-knights/scaffit/pkg/node"
	"github.com/manifoldco/promptui"
	"log"
)

// ModuleGenerator generates the module structure
type ModuleGenerator struct {
	fsys     pkg.FileSystem
	location string
}

// NewModuleGenerator creates a new module generator
func NewModuleGenerator(fsys pkg.FileSystem, location string) *ModuleGenerator {
	return &ModuleGenerator{
		fsys:     fsys,
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
		Validate: golang.ValidateModulePath,
	}

	if modulePath, err = prompt.Run(); err != nil {
		return err
	}

	log.Printf("Module path: %s\n", modulePath)

	if err = golang.InitializeGoModule(g.location, modulePath); err != nil {
		return err
	}

	if err = node.Init(g.location); err != nil {
		return err
	}

	return nil
}

// Generate generates the module structure
func (g *ModuleGenerator) Generate() error {
	var (
		err error
	)

	log.Printf("Generating module at %s\n", g.location)

	if !g.fsys.DirectoryExists(g.location) {
		log.Printf("Creating directory %s\n", g.location)
		if err := g.fsys.CreateDirectory(g.location); err != nil {
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
