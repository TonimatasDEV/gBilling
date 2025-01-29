package smtp

import (
	"log"
	"net/smtp"
	"os"
)

var smtpFullHost = os.Getenv("SMTP_HOST") + ":" + os.Getenv("SMTP_PORT")
var smtpAuth = smtp.PlainAuth(os.Getenv("SMTP_IDENTIFY"), os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
var smtpMail = os.Getenv("SMTP_IDENTIFY")

func SendMail(to []string, subject, body string) {
	message := []byte(subject + "\n" + body)

	err := smtp.SendMail(smtpFullHost, smtpAuth, smtpMail, to, message)
	if err != nil {
		log.Println("Error sending mail:", err)
	}
}
