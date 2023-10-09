package node

import (
	"github.com/jedi-knights/scaffit/pkg"
)

func Init(location string) error {
	var (
		err error
	)

	if err = pkg.RunCommand(location, "npm init --yes", false); err != nil {
		return err
	}
	if err = InstallDependencies(location); err != nil {
		return err
	}

	return nil
}

// InitializeCommitlint initializes commitlint
func InitializeCommitlint(location string) error {
	// What do I have to do to initialize commitlint?
	dependencies := []string{
		"@commitlint/cli",
		"@commitlint/config-conventional",
		"husky",
	}

	for _, dependency := range dependencies {
		if err := InstallDependency(location, dependency); err != nil {
			return err
		}
	}

	if err := pkg.RunCommand(location, `echo "module.exports = { extends: ['@commitlint/config-conventional'] };" > commitlint.config.js`, false); err != nil {
		return err
	}
	if err := pkg.RunCommand(location, "npx husky install", false); err != nil {
		return err
	}
	if err := pkg.RunCommand(location, "npx husky add .husky/commit-msg 'npx --no-install commitlint --edit $1'", false); err != nil {
		return err
	}

	return nil
}

func InstallDependency(location string, dependency string) error {
	if err := pkg.RunCommand(location, "npm install --save-dev --no-fund "+dependency, false); err != nil {
		return err
	}

	return nil
}

func InstallDependencies(location string) error {
	dependencies := []string{
		"@semantic-release/changelog",
		"@semantic-release/commit-analyzer",
		"@semantic-release/exec",
		"@semantic-release/git",
		"@semantic-release/github",
		"@semantic-release/npm",
		"@semantic-release/release-notes-generator",
		"eslint",
		"eslint-config-standard",
		"eslint-plugin-import",
		"eslint-plugin-n",
		"eslint-plugin-promise",
		"semantic-release",
	}

	for _, dependency := range dependencies {
		if err := InstallDependency(location, dependency); err != nil {
			return err
		}
	}

	return nil
}
