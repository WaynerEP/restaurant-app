package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/WaynerEP/restaurant-app/server/plugin/email/global"

	"github.com/jordan-wright/email"
)

// Email
// @function: Email
// @description: Email sending method
// @param: To string, subject string, body string
// @return: error
func Email(To, subject string, body string) error {
	to := strings.Split(To, ",")
	return send(to, subject, body)
}

// ErrorToEmail
// @function: ErrorToEmail
// @description: Send an email to a specified address for email middleware errors
// @param: subject string, body string
// @return: error
func ErrorToEmail(subject string, body string) error {
	to := strings.Split(global.GlobalConfig.To, ",")
	if to[len(to)-1] == "" { // Check if the last element in the slice is empty and remove it
		to = to[:len(to)-1]
	}
	return send(to, subject, body)
}

// EmailTest
// @function: EmailTest
// @description: Email testing method
// @param: subject string, body string
// @return: error
func EmailTest(subject string, body string) error {
	to := []string{global.GlobalConfig.To}
	return send(to, subject, body)
}

// send
// @function: send
// @description: Email sending method
// @param: to []string, subject string, body string
// @return: error
func send(to []string, subject string, body string) error {
	from := global.GlobalConfig.From
	nickname := global.GlobalConfig.Nickname
	secret := global.GlobalConfig.Secret
	host := global.GlobalConfig.Host
	port := global.GlobalConfig.Port
	isSSL := global.GlobalConfig.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}