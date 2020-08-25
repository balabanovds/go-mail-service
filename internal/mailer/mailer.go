package mailer

import (
	"fmt"

	"github.com/jordan-wright/email"
)

type MailSender interface {
	Send(mail Mail) error
}

type Mailer struct {
	config *Config
}

func NewMailer(config *Config) *Mailer {
	return &Mailer{
		config,
	}
}

type Mail struct {
	To      []string
	Subject string
	Body    string
}

func NewMail(to []string, subject string, body string) Mail {
	return Mail{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

func (m *Mailer) Send(mail Mail) error {
	em := email.NewEmail()
	em.From = fmt.Sprintf("%s@%s", m.config.From, m.config.Hostname)
	em.To = mail.To
	em.Subject = mail.Subject
	em.HTML = []byte(mail.Body)

	fmt.Printf("%+v\n", em)

	return em.Send(m.config.Server, nil)

	// c, err := smtp.Dial(m.config.Server)
	// if err != nil {
	// 	return err
	// }
	// defer c.Quit() //nolint:errcheck

	// if err := c.Mail(mail.Sender); err != nil {
	// 	return err
	// }

	// rec := strings.Join(mail.Recipients, ",")
	// if err := c.Rcpt(rec); err != nil {
	// 	return err
	// }

	// wc, err := c.Data()
	// if err != nil {
	// 	return err
	// }

	// body := strings.Join([]string{mail.Subject, mail.Body}, "\n\n")
	// _, err = fmt.Fprint(wc, body)
	// if err != nil {
	// 	return err
	// }

	// if err := wc.Close(); err != nil {
	// 	return err
	// }

	// return nil
}
