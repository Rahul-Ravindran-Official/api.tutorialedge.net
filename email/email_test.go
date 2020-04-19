package email

import "testing"

func TestSendNewUserEmail(t *testing.T) {
	SendNewUserEmail("Unit Testing Email", "Test Body", "testing@tutorialedge.net")
}
