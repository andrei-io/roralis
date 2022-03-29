// Contains everything related to working with emails
package email

import (
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/spf13/viper"
)

type EmailRepo interface {
	Send(recipient string, subject string, email_content string) (*rest.Response, error)
}

type emailRepo struct {
	client *sendgrid.Client
}

// Check interface at compile time
var _ EmailRepo = (*emailRepo)(nil)

// Constructor function
func NewEmailRepo() *emailRepo {
	return &emailRepo{sendgrid.NewSendClient(viper.GetString("SENDGRID_KEY"))}
}

// Sends an email with html formatting
func (e *emailRepo) Send(recipient string, subject string, email_content string) (*rest.Response, error) {

	from := mail.NewEmail(viper.GetString("EMAIL_NAME"), viper.GetString("EMAIL_FROM"))
	to := mail.NewEmail("The User", recipient)
	message := mail.NewSingleEmail(from, subject, to, "", email_content)
	response, err := e.client.Send(message)

	return response, err
}
