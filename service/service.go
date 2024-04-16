package service

import (
	"gitlab.com/nina8884807/mail/entutySend"
	gen "gitlab.com/nina8884807/mail/proto"
	"gopkg.in/gomail.v2"
)

type SendService struct{}

func NewSendService() *SendService {
	return &SendService{}
}

func (s *SendService) SendMessage(msg *gen.SendEmailRequest) error {
	if msg.To == "" {
		return entutySend.ErrNotValidationEmail
	}

	if msg.Text == "" {
		return entutySend.ErrNotValidationText
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "ninamusatova90@gmail.com")
	m.SetHeader("To", msg.To)
	m.SetHeader("Subject", "Account verification")

	m.SetBody("text/plain", msg.Text)

	d := gomail.NewDialer("smtp.gmail.com", 465, "ninamusatova90@gmail.com", "nxbnwzblxgbdsryg")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
