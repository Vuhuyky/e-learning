package smptapp

import (
	"app/internal/connection"
	constant "app/internal/constants"
	logapp "app/pkg/log"
	"net/smtp"
)

func SendEmail(data string, email string) error {
	to := []string{email}
	msg := []byte(data)

	err := smtp.SendMail(
		connection.GetSmtpHost()+":"+connection.GetSmtpPort(),
		connection.GetAuthSmtp(),
		email,
		to,
		msg,
	)
	if err != nil {
		logapp.Logger("send-email", err.Error(), constant.ERROR_LOG)
		return err
	}
	return nil
}
