package pkg

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// FileSystemer is an interface that defines the interactions with the FileSystem
type FileSystemer interface {
	FileExists(path string) bool
	DirectoryExists(path string) bool
	CreateFile(path string) (*os.File, error)
	CreateDirectory(path string, includeGitKeep bool) error
	DeleteFile(path string) error
	DeleteDirectory(path string) error
}

// FileSystem is a struct that implements FileSystemer
//
// FileSystem is used to interact with the file system
type FileSystem struct{}

// NewFileSystem creates a new FileSystem
func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

// FileExists checks if a file exists
func (fs FileSystem) FileExists(path string) bool {
	var (
		err error
	)

	if _, err = os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}

		// Error occurred (e.g., permission denied)
		return false
	}

	return true
}

// DirectoryExists checks if a directory exists
func (fs FileSystem) DirectoryExists(path string) bool {
	// Use os.Stat to check the directory's status
	_, err := os.Stat(path)

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

// CreateFile creates a file
func (fs FileSystem) CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// CreateDirectory creates a directory
//
// uses os.MkdirAll to create the directory and its parent directories if they don't exist
func (fs FileSystem) CreateDirectory(path string, includeGitKeep bool) error {
	var (
		err error
	)

	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	if includeGitKeep {
		gitKeepPath := fmt.Sprintf("%s/.gitkeep", path)
		if _, err = fs.CreateFile(gitKeepPath); err != nil {
			return err
		}
	}

	return nil
}

// DeleteFile deletes a file
func (fs FileSystem) DeleteFile(path string) error {
	var (
		err error
	)

	if err = os.Remove(path); err != nil {
		return err
	}

	return nil
}

// DeleteDirectory deletes a directory
func (fs FileSystem) DeleteDirectory(path string) error {
	var (
		err error
		dir *os.File
	)

	// Open the directory
	if dir, err = os.Open(path); err != nil {
		return err
	}
	defer func(dir *os.File) {
		if err = dir.Close(); err != nil {
			log.Print(err)
		}
	}(dir)

	// Read the directory entries
	entries, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}

	// Delete files and subdirectories
	for _, entry := range entries {
		entryPath := fmt.Sprintf("%s/%s", path, entry)
		if err = os.RemoveAll(entryPath); err != nil {
			return err
		}
	}

	// Remove the directory itself
	if err = os.Remove(path); err != nil {
		return err
	}

	return nil
}

func CopyFilesWithOverwrite(overlayDir, destinationDir string) error {
	return filepath.Walk(overlayDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate the destination path
		destPath := filepath.Join(destinationDir, path[len(overlayDir):])

		// If it's a directory, create it in the destination
		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		// If it's a file, copy it to the destination, overwriting if it exists
		if !info.IsDir() {
			srcFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			destFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
