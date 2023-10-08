package pkg

import (
	"bufio"
	"fmt"
	"os"
)

// ReadVersion reads the version number from the VERSION file
func ReadVersion() (string, error) {
	filename := "VERSION"

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("VERSION file is empty")
}
