package fsys

import (
	"fmt"
	"os"
)

// DirectoryExists checks if a directory exists
func DirectoryExists(dirPath string) bool {
	// Use os.Stat to check the directory's status
	_, err := os.Stat(dirPath)

	if err == nil {
		// Directory exists
		return true
	}

	if os.IsNotExist(err) {
		// Directory does not exist
		return false
	}

	// Error occurred (e.g., permission denied)
	return false
}

// CreateDirectory creates a directory
func CreateDirectory(dirPath string) error {
	// Use os.MkdirAll to create the directory and its parent directories if they don't exist
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDirectory deletes a directory
func DeleteDirectory(dirPath string) error {
	// Open the directory
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	// Read the directory entries
	entries, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}

	// Delete files and subdirectories
	for _, entry := range entries {
		entryPath := fmt.Sprintf("%s/%s", dirPath, entry)
		if err := os.RemoveAll(entryPath); err != nil {
			return err
		}
	}

	// Remove the directory itself
	if err := os.Remove(dirPath); err != nil {
		return err
	}

	return nil
}
