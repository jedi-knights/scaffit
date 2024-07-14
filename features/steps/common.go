package steps

import "github.com/cucumber/godog"

func theExitCodeShouldBe(arg1 int) error {
	return godog.ErrPending
}

func thereShouldBeNoErrors() error {
	return godog.ErrPending
}

func InitializeCommonSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^there should be no errors$`, thereShouldBeNoErrors)
	ctx.Step(`^the exit code should be (\d+)$`, theExitCodeShouldBe)
}
