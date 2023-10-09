package pkg

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// ReadVersion reads the version number from the VERSION file
func ReadVersion() (string, error) {
	filename := "VERSION"

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.Print(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("VERSION file is empty")
}

func DownloadFile(url, filePath string) error {
	// Send an HTTP GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			log.Printf("Error closing response body: %s\n", err)
		}
	}(response.Body)

	// Create or open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.Printf("Error closing file: %s\n", err)
		}
	}(file)

	// Copy the response body to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func RunCommand(location, command string, displayOutput bool) error {
	var (
		err error
	)

	// Command and arguments to run the command
	cmd := exec.Command("sh", "-c", command)

	// Set the working directory for the command (optional)
	cmd.Dir = location

	if displayOutput {
		// Redirect command output to the standard streams
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		// Discard the output
		cmd.Stdout = nil
		cmd.Stderr = nil
	}

	log.Printf("%s\n", command)

	// Run the command
	if err = cmd.Run(); err != nil {
		log.Printf("Command failed: %s\n%v", command, err)

		return err
	}

	return nil
}
