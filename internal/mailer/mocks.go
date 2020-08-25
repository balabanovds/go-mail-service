package mailer

import "errors"

type GoodMailerStub struct{}

func (*GoodMailerStub) Send(mail Mail) error {
	return nil
}

type BadMailerStub struct{}

func (*BadMailerStub) Send(mail Mail) error {
	return errors.New("from bad stub")
}
