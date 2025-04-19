package mail

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/AkifhanIlgaz/hotel-booking-app/config"
	"github.com/go-mail/mail/v2"
)

const sender = "support@vb10.com"

type Email struct {
	From    string
	To      string
	Subject string
	HTML    string
}

type EmailManager struct {
	dialer mail.Dialer
}

func NewManager(config config.SMTPConfig) *EmailManager {
	return &EmailManager{
		dialer: mail.Dialer{
			Host:     config.Host,
			Port:     config.Port,
			Username: config.Username,
			Password: config.Password,
		},
	}
}

func (m *EmailManager) Send(email Email) error {
	msg := mail.NewMessage()

	msg.SetHeader("To", email.To)
	msg.SetHeader("From", sender)
	msg.SetBody("text/html", email.HTML)

	err := m.dialer.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

type resetEmailData struct {
	OTP string
}

func (m *EmailManager) ForgotPassword(to, otp string) error {
	tmpl, err := template.ParseFiles("templates/password_reset.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, resetEmailData{OTP: otp})
	if err != nil {
		return err
	}

	email := Email{
		To:      to,
		Subject: "Reset your password",
		HTML:    body.String(),
	}

	err = m.Send(email)
	if err != nil {
		return fmt.Errorf("forgot password email: %w", err)
	}

	return nil
}
