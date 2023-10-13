package pkg

import "path/filepath"

type Giter interface {
	Commands(location string) []*Command
}

type Git struct {
	fsys FileSystemer
}

func NewGit(fsys FileSystemer) *Git {
	return &Git{
		fsys: fsys,
	}
}

func (g Git) Commands(location string) []*Command {
	var commands []*Command

	gitPath := filepath.Join(location, ".git")

	if !g.fsys.DirectoryExists(gitPath) {
		// The .git directory doesn't exist.
		// Initialize the Git project
		commands = append(commands, NewCommand(location, "git init ."))
	}

	return commands
}
