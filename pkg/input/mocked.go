package input

import "github.com/manifoldco/promptui"

type MockedUserInput struct {
	PromptResponses map[string]string
	SelectResponses map[string]int
	YesNoResponses  map[string]bool
}

func (m MockedUserInput) PromptForString(prompt string, validationFunction promptui.ValidateFunc) (string, error) {
	// Return the mock response for the given prompt
	response, exists := m.PromptResponses[prompt]
	if !exists {
		return "", nil // or an error indicating the prompt was not expected
	}

	// Optionally, call the validator with the response if you want to simulate that behavior
	if validationFunction != nil {
		if err := validationFunction(response); err != nil {
			return "", err
		}
	}

	return response, nil
}

func (m MockedUserInput) SelectFromItems(label string, items []string) (int, string, error) {
	// Return the mock response for the given selection
	index, exists := m.SelectResponses[label]
	if !exists {
		return -1, "", nil // or an error indicating the selection was not expected
	}

	return index, items[index], nil
}

func (m MockedUserInput) SelectFromYesNo(label string) (bool, error) {
	// Define the items corresponding to Yes and No responses

	items := []string{"Yes", "No"}

	// Use SelectFromItems to get the index of the selected item
	index, _, err := m.SelectFromItems(label, items)
	if err != nil {
		return false, err // or handle the error as needed
	}

	// Map the index back to a boolean value
	// Assuming index 0 is "Yes" (true) and index 1 is "No" (false)
	return index == 0, nil
}
