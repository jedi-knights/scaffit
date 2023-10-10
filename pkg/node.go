package pkg

// Noder is an interface that defines the methods for a Node project
type Noder interface {
	Commands(location string) []*Command
	CommitlintCommands(location string) []*Command
	InstallCommand(location, dependency string) *Command
	InstallCommands(location string, dependencies []string) []*Command
	HuskyCommands(location string) []*Command
	CommitlintInstallCommands(location string) []*Command
	EslintInstallCommands(location string) []*Command
	SemanticReleaseInstallCommands(location string) []*Command
}

// Node is a struct that implements the Noder interface
type Node struct {
	useFlags map[string]bool
}

// NewNode creates a new Node instance
func NewNode(useFlags map[string]bool) *Node {
	return &Node{
		useFlags: useFlags,
	}
}

// Commands returns a slice of commands to initialize a project
//
// This function uses the useFlags eslint, commitlint, and semantic-release.
func (n Node) Commands(location string) []*Command {
	var commands []*Command

	// Initialize the Node project
	commands = append(commands, NewCommand(location, "npm init --yes"))

	if n.useFlags["semantic-release"] {
		// Install Semantic Release dependencies
		commands = append(commands, n.SemanticReleaseInstallCommands(location)...)
	}

	if n.useFlags["eslint"] {
		commands = append(commands, n.EslintInstallCommands(location)...)
	}

	if n.useFlags["commitlint"] {
		// Install Commitlint dependencies
		commands = append(commands, n.CommitlintInstallCommands(location)...)

		// Install Husky
		commands = append(commands, n.HuskyCommands(location)...)

		// Create the commitlint config file
		commands = append(commands, NewCommand(location, "echo \"module.exports = { extends: ['@commitlint/config-conventional'] };\" > commitlint.config.js"))
	}

	return commands
}

// HuskyCommands returns a slice of commands to initialize a project
func (n Node) HuskyCommands(location string) []*Command {
	return []*Command{
		n.InstallCommand(location, "husky"),
		NewCommand(location, "npx husky install"),
		NewCommand(location, "npx husky add .husky/commit-msg 'npx --no-install commitlint --edit $1'"),
	}
}

// CommitlintInstallCommands returns a slice of commands to initialize a project
func (n Node) CommitlintInstallCommands(location string) []*Command {
	return n.InstallCommands(location, []string{
		"@commitlint/cli",
		"@commitlint/config-conventional",
	})
}

// EslintInstallCommands returns a slice of commands to initialize a project
func (n Node) EslintInstallCommands(location string) []*Command {
	return n.InstallCommands(location, []string{
		"eslint",
		"eslint-config-standard",
		"eslint-plugin-import",
		"eslint-plugin-n",
		"eslint-plugin-promise",
	})
}

// SemanticReleaseInstallCommands returns a slice of commands to initialize a project
func (n Node) SemanticReleaseInstallCommands(location string) []*Command {
	return n.InstallCommands(location, []string{
		"@semantic-release/changelog",
		"@semantic-release/commit-analyzer",
		"@semantic-release/exec",
		"@semantic-release/git",
		"@semantic-release/github",
		"@semantic-release/npm",
		"@semantic-release/release-notes-generator",
		"semantic-release",
	})
}

// InstallCommand returns a command to install a dependency
func (n Node) InstallCommand(location, dependency string) *Command {
	return NewCommand(location, "npm install --save-dev --no-fund "+dependency)
}

// InstallCommands returns a slice of commands to install the dependencies
func (n Node) InstallCommands(location string, dependencies []string) []*Command {
	var commands []*Command

	for _, dependency := range dependencies {
		commands = append(commands, n.InstallCommand(location, dependency))
	}

	return commands
}
