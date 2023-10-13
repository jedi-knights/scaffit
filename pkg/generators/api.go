package generators

import (
	"github.com/jedi-knights/scaffit/pkg"
	"github.com/manifoldco/promptui"
	"log"
	"path/filepath"
	"strings"
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
		Label: "Which framework would you like to use?",
		Items: []string{"Echo", "Gin", "Fiber", "None"},
	}

	if _, answer, err = selectUi.Run(); err != nil {
		return err
	}

	switch answer {
	case "Gin":
		useFlags["gin"] = true
		useFlags["echo"] = false
		useFlags["fiber"] = false
		useFlags["none"] = false
	case "Echo":
		useFlags["gin"] = false
		useFlags["echo"] = true
		useFlags["fiber"] = false
		useFlags["none"] = false
	case "Fiber":
		useFlags["gin"] = false
		useFlags["echo"] = false
		useFlags["fiber"] = true
		useFlags["none"] = false
	case "None":
		useFlags["gin"] = false
		useFlags["echo"] = false
		useFlags["fiber"] = false
		useFlags["none"] = true
	default:
		useFlags["gin"] = false
		useFlags["echo"] = false
		useFlags["fiber"] = false
		useFlags["none"] = true
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
	commands = append(commands, pkg.NewGit(g.fsys).Commands(localModulePath)...)
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

	if useFlags["gin"] {
		return g.generateGinApi(localModulePath, useFlags)
	} else if useFlags["echo"] {
		return g.generateEchoApi(localModulePath, useFlags)
	} else if useFlags["fiber"] {
		return g.generateFiberApi(localModulePath, useFlags)
	} else {
		// Don't do anything extra.
		log.Printf("No framework selected\n")
	}

	return nil
}

// generateGinApi generates the gin api
func (g *ApiGenerator) generateGinApi(location string, useFlags map[string]bool) error {
	log.Println("Generating gin api ...")
	return nil
}

// generateEchoApi generates the echo api
func (g *ApiGenerator) generateEchoApi(location string, useFlags map[string]bool) error {
	log.Println("Generating echo api ...")

	commands := []*pkg.Command{}

	golang := pkg.NewGolang(useFlags)

	// Install echo specific dependencies
	commands = append(commands, golang.InstallCommand(location, "github.com/labstack/echo/v4"))
	commands = append(commands, golang.InstallCommand(location, "github.com/swaggo/echo-swagger"))
	commands = append(commands, golang.InstallCommand(location, "github.com/swaggo/swag"))
	commands = append(commands, pkg.NewCommand(location, "mkdir certs"))

	//// Relative path to the openssl.cnf file
	//relativePath := "openssl.cnf"
	//
	//// Get the absolute path based on the current working directory
	//absPath, err := filepath.Abs(relativePath)
	//if err != nil {
	//	panic(err)
	//}

	// commands = append(commands, pkg.NewCommand(location, fmt.Sprintf("openssl req -x509 -nodes -newkey rsa:2048 -keyout certs/key.pem -out certs/cert.pem -days 365 -config %s", absPath)))
	// commands = append(commands, pkg.NewCommand(location, "openssl req -x509 -nodes -newkey rsa:2048 -keyout certs/key.pem -out certs/cert.pem -days 365"))
	commands = append(commands, pkg.NewCommand(location, "touch certs/.gitkeep"))
	commands = append(commands, pkg.NewCommand(location, "mkdir data"))
	commands = append(commands, pkg.NewCommand(location, "touch data/.gitkeep"))
	commands = append(commands, pkg.NewCommand(location, "mkdir db"))
	commands = append(commands, pkg.NewCommand(location, "touch db/.gitkeep"))
	commands = append(commands, pkg.NewCommand(location, "mkdir handlers"))
	commands = append(commands, pkg.NewCommand(location, "touch handlers/.gitkeep"))
	commands = append(commands, pkg.NewCommand(location, "mkdir types"))
	commands = append(commands, pkg.NewCommand(location, "touch types/.gitkeep"))
	commands = append(commands, pkg.NewCommand(location, "touch .env"))

	if useFlags["cobra"] {
		commands = append(commands, pkg.NewCommand(location, "cobra-cli add serve"))
		commands = append(commands, pkg.NewCommand(location, "cobra-cli add seed"))
		commands = append(commands, pkg.NewCommand(location, "cobra-cli add migrate"))
		commands = append(commands, pkg.NewCommand(location, "cobra-cli add drop"))
	}

	// Todo: Create default handler for version
	// Todo: Create default handler for health
	// Todo: Create default handler for todos
	// The todos handlers will contain basic CRUD examples for easy access.

	for _, command := range commands {
		if strings.Contains(command.Value, "openssl") {
			if err := command.Execute(true); err != nil {
				return err
			}
		} else {
			if err := command.Execute(false); err != nil {
				return err
			}
		}
	}

	// Overlay the files from data/api/echo/overlay into the location
	log.Println("Copying overlay files ...")
	if err := pkg.CopyFilesWithOverwrite("data/api/echo/overlay", location); err != nil {
		return err
	}

	return nil
}

// generateFiberApi generates the fiber api
func (g *ApiGenerator) generateFiberApi(location string, useFlags map[string]bool) error {
	log.Println("Generating fiber api ...")
	return nil
}

// Generate generates the api structure
func (g *ApiGenerator) Generate() error {
	var (
		err error
	)

	log.Printf("Generating api at %s\n", g.location)

	if !g.fsys.DirectoryExists(g.location) {
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
