package features

import (
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/jedi-knights/scaffit/features/steps"
	"os"
	"testing"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	steps.InitializeAPISteps(ctx)
	steps.InitializeCLISteps(ctx)
	steps.InitializeCommonSteps(ctx)
	steps.InitializeModuleSteps(ctx)
	steps.InitializeNewSteps(ctx)
	steps.InitializeVersionSteps(ctx)
}

func TestMain(m *testing.M) {
	status := godog.TestSuite{
		Name:                 "acceptance",
		TestSuiteInitializer: func(ctx *godog.TestSuiteContext) {},
		ScenarioInitializer:  InitializeScenario,
		Options: &godog.Options{
			Format:    "pretty",
			Output:    colors.Colored(os.Stdout),
			Paths:     []string{"features"},
			Randomize: 0,
		},
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
