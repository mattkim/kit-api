package email

import (
	"log"
	"net/smtp"
	"os"
)

// Send ...
func Send(subject string, body string, to string) {
	// TODO: add these to configs
	from := os.Getenv("KIT_EMAIL")
	password := os.Getenv("KIT_EMAIL_PASSWORD")

	msg := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"MIME-Version: 1.0" + "\r\n" +
		"Content-type: text/html" + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body + "\r\n"

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, password, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	log.Print("message sent")
}
