package golang

import (
	"os"
	"os/exec"
)

func InitializeGoModule(location, modulePath string) error {
	// Command and arguments to run "go mod init" with the specified module path
	cmd := exec.Command("go", "mod", "init", modulePath)

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
