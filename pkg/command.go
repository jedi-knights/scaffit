package pkg

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Commder interface {
	Execute() error
}

type Command struct {
	Location      string
	Value         string
	DisplayOutput bool
}

func NewCommand(location string, value string) *Command {
	return &Command{
		Location: location,
		Value:    value,
	}
}

// Execute runs the command
func (c *Command) Execute(displayOutput bool) error {
	var (
		err error
	)

	// Command and arguments to run the command
	cmd := exec.Command("sh", "-c", c.Value)

	// Set the working directory for the command (optional)
	cmd.Dir = c.Location

	if displayOutput {
		// Redirect command output to the standard streams
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		// Discard the output
		cmd.Stdout = nil
		cmd.Stderr = nil
	}

	log.Printf("%s\n", c.Value)

	// Run the command
	if err = cmd.Run(); err != nil {
		log.Printf("Command failed: %s\n%v", c.Value, err)

		return err
	}

	return nil
}

func (c *Command) String() string {
	return fmt.Sprintf("Command{Location: '%s', Value: '%s'}", c.Location, c.Value)
}
