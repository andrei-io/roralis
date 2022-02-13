// Contains everything related to working with emails
package email

import (
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/spf13/viper"
)

type IEmailRepo interface {
	Send(recipient string, subject string, email_content string) (*rest.Response, error)
}

type EmailRepo struct {
	client *sendgrid.Client
}

// Constructor function
func NewEmailRepo() *EmailRepo {
	return &EmailRepo{sendgrid.NewSendClient(viper.GetString("SENDGRID_KEY"))}
}

// Sends an email with html formatting
func (e *EmailRepo) Send(recipient string, subject string, email_content string) (*rest.Response, error) {

	from := mail.NewEmail(viper.GetString("EMAIL_NAME"), viper.GetString("EMAIL_FROM"))
	to := mail.NewEmail("The User", recipient)
	message := mail.NewSingleEmail(from, subject, to, "", email_content)
	response, err := e.client.Send(message)

	return response, err
}
