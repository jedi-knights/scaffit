package pkg

type Giter interface {
	Commands(location string) []*Command
}

type Git struct{}

func NewGit() *Git {
	return &Git{}
}

func (g Git) Commands(location string) []*Command {
	var commands []*Command

	// Initialize the Git project
	commands = append(commands, NewCommand(location, "git init ."))

	return commands
}
