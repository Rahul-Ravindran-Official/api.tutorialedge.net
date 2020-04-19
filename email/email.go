package email

import (
	"fmt"
	"os"

	"github.com/mailgun/mailgun-go"
)

// SendNewUserEmail - Sends a notification to the TutorialEdge email group notifying a new user
//  has signed up to the site.
func SendNewUserEmail(subject, body, recipient string) error {
	fmt.Println("Sending New User Email Notification...")
	
	mailgunAPIKey := os.Getenv("MAILGUN_API_KEY")
	mg := mailgun.NewMailgun("tutorialedge.net", mailgunAPIKey)
	m := mg.NewMessage(
		"Notifications <admin@tutorialedge.net>",
		subject,
		body,
		recipient,
	)
	resp, id, err := mg.Send(m)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("ID: %s\n", id)
	fmt.Printf("Response: %s\n", resp)

	return nil
}
