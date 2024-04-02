package new_service

import (
	gen "gitlab.com/nina8884807/mail/proto"
	"gopkg.in/gomail.v2"
)

func SendMessage(msg *gen.SendEmailRequest) error {

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
