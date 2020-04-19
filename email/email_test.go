// +build integration

package email

import "testing"

func TestSendEmail(t *testing.T) {
	err := SendEmail("Unit Testing Email", "Test Body", "testing@tutorialedge.net")
	if err != nil {
		t.Error(err)
	}
}
