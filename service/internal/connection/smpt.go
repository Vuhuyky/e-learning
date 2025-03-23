package connection

import (
	"net/smtp"
)

func initSmptAuth() {
	authSmtp = smtp.PlainAuth(
		"",
		conn.Smpt.Email,
		conn.Smpt.Password,
		conn.Smpt.Host,
	)
}
