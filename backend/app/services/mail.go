package services

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/resend/resend-go/v2"
)

var client = resend.NewClient(os.Getenv("RESEND_API_KEY"))

// TODO add role and team name
// TODO add HTML template
func SendTeamInvitationEmail(recipientEmail string) error {
	data := struct {
		Email            string
		InviterFirstName string
		InviterEmail     string
		TeamName         string
		InvitationLink   string
	}{
		Email:            recipientEmail,
		InviterFirstName: "Veronica",
		InviterEmail:     "veronica@poised.com",
		TeamName:         "Poised",
		InvitationLink:   "https://app.lockform.io",
	}
	// Read the HTML template from a file
	tmpl, err := template.ParseFiles("assets/templates/emails/lockform-user-invite.html")
	if err != nil {
		fmt.Println("Failed to read template file:", err)
		return err
	}
	var renderedTemplate strings.Builder
	err = tmpl.Execute(&renderedTemplate, data)
	if err != nil {
		fmt.Println("Failed to render template:", err)
		return err
	}
	params := &resend.SendEmailRequest{
		From:    "Lockform <onboarding@rentisha.com>",
		To:      []string{recipientEmail},
		Subject: fmt.Sprintf("Invitation to Join %s on Lockform", data.TeamName),
		Html:    renderedTemplate.String(),
	}

	if _, err := client.Emails.Send(params); err != nil {
		return err
	}
	return nil
}
