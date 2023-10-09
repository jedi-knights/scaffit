package generators

import (
	"github.com/jedi-knights/scaffit/pkg"
	"github.com/jedi-knights/scaffit/pkg/golang"
	"github.com/jedi-knights/scaffit/pkg/node"
	"github.com/manifoldco/promptui"
	"log"
	"path/filepath"
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
		err           error
		answer        string
		modulePath    string
		useCobra      bool
		useViper      bool
		useCommitlint bool
		selectUi      promptui.Select
		promptUi      promptui.Prompt
	)

	promptUi = promptui.Prompt{
		Label:    "Module path",
		Validate: golang.ValidateModulePath,
	}

	if modulePath, err = promptUi.Run(); err != nil {
		return err
	}

	selectUi = promptui.Select{
		Label: "Would you like to use Cobra/Viper?",
		Items: []string{"Cobra and Viper", "Cobra Only", "Neither"},
	}

	if _, answer, err = selectUi.Run(); err != nil {
		return err
	}

	switch answer {
	case "Cobra Only":
		useCobra = true
		useViper = false
	case "Cobra and Viper":
		useCobra = true
		useViper = true
	case "Neither":
		useCobra = false
		useViper = false
	default:
		useCobra = false
		useViper = false
	}

	selectUi = promptui.Select{
		Label: "Would you like to use conventional commits?",
		Items: []string{"Yes", "No"},
	}

	if _, answer, err = selectUi.Run(); err != nil {
		return err
	}

	useCommitlint = answer == "Yes"

	log.Printf("Module path: %s\n", modulePath)

	// We need to build the local module path here.
	localModulePath := filepath.Join(g.location, modulePath)

	log.Printf("Local module path: %s\n", localModulePath)

	// Create the local module path
	if err = g.fsys.CreateDirectory(localModulePath, false); err != nil {
		return err
	}

	// Initialize the Git repository
	if err = pkg.RunCommand(localModulePath, "git init .", false); err != nil {
		return err
	}

	if err = golang.InitializeGoModule(localModulePath, modulePath); err != nil {
		return err
	}

	if useCobra {
		// Using Cobra
		log.Printf("Using Cobra\n")
		if err = golang.InitializeCobra(localModulePath, useViper); err != nil {
			return err
		}
	} else {
		// Not using Cobra
		log.Printf("Not using Cobra\n")
		if useViper {
			log.Fatal("Cannot use Viper without Cobra at this time.")
		}
	}

	if err = node.Init(localModulePath); err != nil {
		return err
	}

	if useCommitlint {
		log.Printf("Using conventional commits\n")
		if err = node.InitializeCommitlint(localModulePath); err != nil {
			return err
		}
	} else {
		log.Printf("Not using conventional commits\n")
	}

	// Download a markdown file
	uri := "https://raw.githubusercontent.com/Ismaestro/markdown-template/master/README.md"
	readmePath := filepath.Join(localModulePath, "README.md")
	if err = pkg.DownloadFile(uri, readmePath); err != nil {
		return err
	}

	// Create pkg directory
	pkgPath := filepath.Join(localModulePath, "pkg")
	if err = g.fsys.CreateDirectory(pkgPath, true); err != nil {
		return err
	}

	// Create utils directory
	utilsPath := filepath.Join(localModulePath, "utils")
	if err = g.fsys.CreateDirectory(utilsPath, true); err != nil {
		return err
	}

	// Create an assets directory
	assetsPath := filepath.Join(localModulePath, "assets")
	if err = g.fsys.CreateDirectory(assetsPath, true); err != nil {
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
		if err := g.fsys.CreateDirectory(g.location, false); err != nil {
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
