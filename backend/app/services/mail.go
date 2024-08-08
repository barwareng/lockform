package services

import (
	"os"

	"github.com/resend/resend-go/v2"
)

var client = resend.NewClient(os.Getenv("RESEND_API_KEY"))

// TODO add role and team name
// TODO add HTML template
func SendTeamInvitationEmail(recipientEmail string) error {
	params := &resend.SendEmailRequest{
		From:    "support@rentisha.com",
		To:      []string{recipientEmail},
		Subject: "Welcome to Team",
		Html:    "<p>Congrats on sending your <strong>first email</strong>!</p>",
	}

	if _, err := client.Emails.Send(params); err != nil {
		return err
	}
	return nil
}
