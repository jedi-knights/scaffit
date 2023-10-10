package generators

import (
	"github.com/jedi-knights/scaffit/pkg"
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
		err        error
		answer     string
		modulePath string
		selectUi   promptui.Select
		promptUi   promptui.Prompt
		useFlags   map[string]bool
		commands   []*pkg.Command
	)

	useFlags = make(map[string]bool)

	promptUi = promptui.Prompt{
		Label:    "Module path",
		Validate: pkg.ValidateModulePath,
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
		useFlags["cobra"] = true
		useFlags["viper"] = false
	case "Cobra and Viper":
		useFlags["cobra"] = true
		useFlags["viper"] = true
	case "Neither":
		useFlags["cobra"] = false
		useFlags["viper"] = false
	default:
		useFlags["cobra"] = false
		useFlags["viper"] = false
	}

	selectUi = promptui.Select{
		Label: "Would you like to use conventional commits?",
		Items: []string{"Yes", "No"},
	}
	if _, answer, err = selectUi.Run(); err != nil {
		return err
	}
	useFlags["commitlint"] = answer == "Yes"

	selectUi = promptui.Select{
		Label: "Would you like to use semantic release?",
		Items: []string{"Yes", "No"},
	}
	if _, answer, err = selectUi.Run(); err != nil {
		return err
	}
	useFlags["semantic-release"] = answer == "Yes"

	selectUi = promptui.Select{
		Label: "Would you like to use eslint?",
		Items: []string{"Yes", "No"},
	}
	if _, answer, err = selectUi.Run(); err != nil {
		return err
	}
	useFlags["eslint"] = answer == "Yes"

	log.Printf("Module path: %s\n", modulePath)

	// We need to build the local module path here.
	localModulePath := filepath.Join(g.location, modulePath)

	log.Printf("Local module path: %s\n", localModulePath)

	commands = append(commands, pkg.NewCommand(g.location, "mkdir -p "+modulePath))
	commands = append(commands, pkg.NewGit().Commands(localModulePath)...)
	commands = append(commands, pkg.NewNode(useFlags).Commands(localModulePath)...)
	commands = append(commands, pkg.NewGolang(useFlags).Commands(localModulePath, modulePath)...)

	commands = append(commands, pkg.NewCommand(localModulePath, "curl -o README.md https://raw.githubusercontent.com/Ismaestro/markdown-template/master/README.md"))

	commands = append(commands, pkg.NewCommand(localModulePath, "mkdir pkg"))
	commands = append(commands, pkg.NewCommand(localModulePath, "echo 'package pkg' > pkg/pkg.go"))
	commands = append(commands, pkg.NewCommand(localModulePath, "mkdir utils"))
	commands = append(commands, pkg.NewCommand(localModulePath, "echo 'package utils' > utils/utils.go"))
	commands = append(commands, pkg.NewCommand(localModulePath, "mkdir assets"))
	commands = append(commands, pkg.NewCommand(localModulePath, "touch assets/.gitkeep"))

	for _, command := range commands {
		if err = command.Execute(false); err != nil {
			return err
		}
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
