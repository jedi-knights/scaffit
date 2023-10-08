package fsys

import "os"

// FileExists checks if a file exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		// File exists
		return true
	}
	if os.IsNotExist(err) {
		// File does not exist
		return false
	}
	// Error occurred (e.g., permission denied)
	return false
}

// CreateFile creates a file
func CreateFile(filePath string) (*os.File, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// DeleteFile deletes a file
func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
