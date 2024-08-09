package services

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/lockform/app/models"
	"github.com/resend/resend-go/v2"
)

var client = resend.NewClient(os.Getenv("RESEND_API_KEY"))

func SendEmailVerificationLink(email string, link string) error {

	// Read the HTML template from a file
	tmpl, err := template.ParseFiles("assets/templates/emails/email-verification.html")
	if err != nil {
		fmt.Println("Failed to read template file:", err)
		return err
	}
	var renderedTemplate strings.Builder
	data := struct {
		VerificationLink string
	}{

		VerificationLink: link,
	}
	err = tmpl.Execute(&renderedTemplate, data)
	if err != nil {
		fmt.Println("Failed to render template:", err)
		return err
	}
	params := &resend.SendEmailRequest{
		From:    "Lockform <onboarding@rentisha.com>",
		To:      []string{email},
		Subject: "Verify your email address.",
		Html:    renderedTemplate.String(),
	}

	if _, err := client.Emails.Send(params); err != nil {
		return err
	}
	return nil
}
func SendPasswordResetLink(email string, link string) error {
	// Read the HTML template from a file
	tmpl, err := template.ParseFiles("assets/templates/emails/password-reset.html")
	if err != nil {
		fmt.Println("Failed to read template file:", err)
		return err
	}
	var renderedTemplate strings.Builder
	data := struct {
		ResetLink string
	}{

		ResetLink: link,
	}
	err = tmpl.Execute(&renderedTemplate, data)
	if err != nil {
		fmt.Println("Failed to render template:", err)
		return err
	}
	params := &resend.SendEmailRequest{
		From:    "Lockform <onboarding@rentisha.com>",
		To:      []string{email},
		Subject: "Reset your password.",
		Html:    renderedTemplate.String(),
	}

	if _, err := client.Emails.Send(params); err != nil {
		return err
	}
	return nil
}
func SendTeamInvitationEmail(data models.InviteTeamUser) error {

	// Read the HTML template from a file
	tmpl, err := template.ParseFiles("assets/templates/emails/user-invitation.html")
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
		To:      []string{data.Email},
		Subject: fmt.Sprintf("Invitation to Join %s on Lockform", data.TeamName),
		Html:    renderedTemplate.String(),
	}

	if _, err := client.Emails.Send(params); err != nil {
		return err
	}
	return nil
}
