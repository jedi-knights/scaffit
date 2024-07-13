package pkg

import (
	"io"
	"os"
	"strings"
)

//// ReadVersion reads the version number from the VERSION file
//func ReadVersion() (string, error) {
//	filename := "VERSION"
//
//	file, err := os.Open(filename)
//	if err != nil {
//		return "", err
//	}
//	defer func(file *os.File) {
//		if err = file.Close(); err != nil {
//			log.Print(err)
//		}
//	}(file)
//
//	scanner := bufio.NewScanner(file)
//	if scanner.Scan() {
//		return scanner.Text(), nil
//	}
//
//	if err := scanner.Err(); err != nil {
//		return "", err
//	}
//
//	return "", fmt.Errorf("VERSION file is empty")
//}

func ReadVersion() (string, error) {
	// Open the VERSION file
	file, err := os.Open("VERSION")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the content of the file
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Convert the content to a string and trip white spaces
	version := string(content)
	version = strings.TrimSpace(version)

	return version, nil
}
