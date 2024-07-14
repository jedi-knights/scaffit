package steps

import (
	"github.com/cucumber/godog"
	"os/exec"
)

var lastCommandOutput string
var lastCommandError error

func iExecuteTheVersionCommand() error {
	cmd := exec.Command("scaffit", "version")
	out, err := cmd.CombinedOutput()
	lastCommandOutput = string(out)
	lastCommandError = err

	return err
}

func iShouldSeeTheVersionFromTheVERSIONFile() error {
	return godog.ErrPending
}

func InitializeVersionSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I execute the version command$`, iExecuteTheVersionCommand)
	ctx.Step(`^I should see the version from the VERSION file$`, iShouldSeeTheVersionFromTheVERSIONFile)
}
