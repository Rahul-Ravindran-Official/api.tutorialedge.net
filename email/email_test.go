package email

import "testing"

func TestSendNewUserEmail(t *testing.T) {
	err := SendNewUserEmail("Unit Testing Email", "Test Body", "testing@tutorialedge.net")
	if err != nil {
		t.Error(err)
	}
}
