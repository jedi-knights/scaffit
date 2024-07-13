package pkg

import (
	"fmt"
	"log"
)

const ignoreListUri = "https://raw.githubusercontent.com/github/gitignore/master/Go.gitignore"

type Golanger interface {
	Commands(location, modulePath string) ([]*Command, error)
}

type Golang struct {
	useFlags map[string]bool
}

func NewGolang(useFlags map[string]bool) *Golang {
	return &Golang{
		useFlags: useFlags,
	}
}

func (g Golang) InstallCommand(location, dependency string) *Command {
	return NewCommand(location, "go get "+dependency)
}

func (g Golang) InstallCommands(location string, dependencies []string) []*Command {
	var commands []*Command

	for _, dependency := range dependencies {
		commands = append(commands, g.InstallCommand(location, dependency))
	}

	return commands
}

func (g Golang) Commands(location, modulePath string) []*Command {
	var (
		commands []*Command
	)

	// Initialize the Go module
	commands = append(commands, NewCommand(location, "go mod init "+modulePath))
	commands = append(commands, NewCommand(location, fmt.Sprintf("curl -o .gitignore %s", ignoreListUri)))

	if g.useFlags["cobra"] {
		commands = append(commands, g.InstallCommand(location, "github.com/spf13/cobra"))

		if g.useFlags["viper"] {
			commands = append(commands, g.InstallCommand(location, "github.com/spf13/viper"))

			log.Print("Using Cobra and Viper\n")
			commands = append(commands, NewCommand(location, "cobra-cli init --viper"))
		} else {
			log.Print("Using Cobra only\n")
			commands = append(commands, NewCommand(location, "cobra-cli init"))
		}
	} else {
		if g.useFlags["viper"] {
			log.Print("Using Viper only\n")
			commands = append(commands, g.InstallCommand(location, "github.com/spf13/viper"))

			log.Print("Todo: Figure out what we need to do to setup viper only.")
		} else {
			log.Print("Not using Cobra or Viper\n")
		}
	}

	return commands
}
