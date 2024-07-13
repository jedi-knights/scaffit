package features

import (
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/jedi-knights/scaffit/features/steps" // Replace "your_project/test" with the actual import path
	"os"
	"testing"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	steps.InitializeScenario(ctx) // Assuming `InitializeScenario` is the function in your `test` package that initializes steps
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
