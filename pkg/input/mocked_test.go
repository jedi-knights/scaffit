package input

import "testing"

func TestMockedUserInput_PromptForString(t *testing.T) {
	mockInput := MockedUserInput{
		PromptResponses: map[string]string{
			"test prompt": "expected response",
		},
	}

	got, err := mockInput.PromptForString("test prompt", nil)
	if err != nil {
		t.Errorf("PromptForString() error = %v", err)
	}
	if got != "expected response" {
		t.Errorf("PromptForString() = %v, want %v", got, "expected response")
	}
}

func TestMockedUserInput_SelectFromItems(t *testing.T) {
	mockInput := MockedUserInput{
		SelectResponses: map[string]int{
			"test select": 1,
		},
	}

	_, got, err := mockInput.SelectFromItems("test select", []string{"Option 1", "Option 2"})
	if err != nil {
		t.Errorf("SelectFromItems() error = %v", err)
	}
	if got != "Option 2" {
		t.Errorf("SelectFromItems() = %v, want %v", got, "Option 2")
	}
}

func TestMockedUserInput_SelectFromYesNo(t *testing.T) {
	mockInput := MockedUserInput{
		SelectResponses: map[string]int{
			"test yes/no": 0, // Assuming 0 is "Yes"
		},
	}

	got, err := mockInput.SelectFromYesNo("test yes/no")
	if err != nil {
		t.Errorf("SelectFromYesNo() error = %v", err)
	}
	if !got {
		t.Errorf("SelectFromYesNo() = %v, want %v", got, true)
	}
}
