package service

import (
	"github.com/WaynerEP/restaurant-app/server/plugin/email/utils"
)

type EmailService struct{}

// EmailTest
// @function: EmailTest
// @description: Send a test email
// @return: err error
func (e *EmailService) EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = utils.EmailTest(subject, body)
	return err
}

// SendEmail
// @function: SendEmail
// @description: Send an email
// @return: err error
// @param to string      Recipient
// @param subject string Subject (Title)
// @param body  string      Email content
func (e *EmailService) SendEmail(to, subject, body string) (err error) {
	err = utils.Email(to, subject, body)
	return err
}
