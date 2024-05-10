package service

import (
	"gitlab.com/nina8884807/mail/entutySend"
	gen "gitlab.com/nina8884807/mail/proto"
	"gopkg.in/gomail.v2"
)

type SendService struct {
	mailLogin    string
	mailPassword string
}

func NewSendService(mailLogin string, mailPassword string) *SendService {
	return &SendService{mailLogin: mailLogin, mailPassword: mailPassword}
}

func (s *SendService) SendMessage(msg *gen.SendEmailRequest) error {
	if msg.To == "" {
		return entutySend.ErrNotValidationEmail
	}

	if msg.Text == "" {
		return entutySend.ErrNotValidationText
	}

	m := gomail.NewMessage()
	m.SetHeader("From", s.mailLogin)
	m.SetHeader("To", msg.To)
	m.SetHeader("Subject", msg.Subject)

	m.SetBody("text/plain", msg.Text)

	d := gomail.NewDialer("smtp.gmail.com", 465, s.mailLogin, s.mailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
