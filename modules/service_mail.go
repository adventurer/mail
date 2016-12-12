package service

import (
	"fmt"
	"net/mail"
	"net/smtp"

	"github.com/scorredoira/email"
)

type Service struct {
}

func (this *Service) SendToMail(host, from, to, pass, title, content, typen string) bool {
	fmt.Println("sending...")
	// compose the message
	m := new(email.Message)
	if typen == "html" {
		m = email.NewHTMLMessage(title, content)
	} else {
		m = email.NewMessage(title, content)
	}

	m.From = mail.Address{Name: from, Address: from}
	m.To = []string{to}

	// add attachments
	// if err := m.Attach("email.go"); err != nil {
	// 	log.Fatal(err)
	// }

	// send it
	auth := smtp.PlainAuth("", from, pass, host)
	if err := email.Send(host+":25", auth, m); err != nil {
		fmt.Println("big error!!!!")
		fmt.Println(err)
		return false
	}

	return true
}
