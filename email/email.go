package email

import (
	"fmt"
	"os"

	"github.com/mailgun/mailgun-go"
)

// SendNewUserEmail - Sends a notification to the TutorialEdge email group notifying a new user
//  has signed up to the site.
func SendNewUserEmail(body string) {
	fmt.Println("Sending New User Email Notification...")

	mailgunAPIKey := os.Getenv("MAILGUN_API_KEY")
	mg := mailgun.NewMailgun("tutorialedge.net", mailgunAPIKey)
	m := mg.NewMessage(
		"Notifications <admin@tutorialedge.net>",
		"New User Registration",
		body,
		"elliot@tutorialedge.net",
	)
	_, id, err := mg.Send(m)
	fmt.Printf("ID: %s\n", id)

	if err != nil {
		fmt.Println(err)
	}
}
