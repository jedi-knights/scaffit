package input

import "github.com/manifoldco/promptui"

type RealUserInput struct{}

func (r RealUserInput) PromptForString(label string, validationFunction promptui.ValidateFunc) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validationFunction,
	}

	return prompt.Run()
}

func (r RealUserInput) SelectFromItems(label string, items []string) (int, string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	return prompt.Run()
}

func (r RealUserInput) SelectFromYesNo(label string) (bool, error) {
	prompt := promptui.Select{
		Label: label,
		Items: []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()

	return result == "Yes", err
}
