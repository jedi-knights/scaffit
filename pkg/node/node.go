package node

import (
	"os"
	"os/exec"
)

func Init(location string) error {
	// Command and arguments to run npm init with --yes flag
	cmd := exec.Command("npm", "init", "--yes")

	// Set the working directory for the command (optional)
	cmd.Dir = location

	// Redirect command output to the standard streams
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
