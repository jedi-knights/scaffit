package input

import "github.com/manifoldco/promptui"

type UserInput interface {
	PromptForString(label string, validationFunction promptui.ValidateFunc) (string, error)
	SelectFromItems(label string, items []string) (int, string, error)
	SelectFromYesNo(label string) (bool, error)
}
