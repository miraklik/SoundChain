package utils

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"soundchain/config/config"
)

func SendEmail(fromEmail, toEmail string) error {
	cfg, err := config.Load()
	if err != nil {
		log.Printf("Failed to load config: %v", err)
		return errors.New("failed to load config")
	}

	from := fromEmail
	to := toEmail

	subject := "Confirmation letter"
	body := "Your activate code is " + GenerateRandomString()
	message := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", from, to, subject, body)

	auth := smtp.PlainAuth("", cfg.Email.SmtpUser, cfg.Email.SmtpPass, cfg.Email.SmtpHOST)

	log.Printf("Attempting to send email to %s via %s:%s", to, cfg.Email.SmtpHOST, cfg.Email.SmtpPort)
	err = smtp.SendMail(cfg.Email.SmtpHOST+":"+cfg.Email.SmtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Printf("Failed to send email to %s: %v", to, err)
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
