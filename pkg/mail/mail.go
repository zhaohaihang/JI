package mail

import (
	"ji/config"
	"net/smtp"
	"net/textproto"

	"github.com/google/wire"
	"github.com/jordan-wright/email"
)

const (
	POOL_SIZE  = 4
	EMAIL_FROM = "JI office<1932859223@qq.com>"
)

type MailClient struct {
	mailPool *email.Pool
}

func NewMailClient(cfg *config.Config) (*MailClient, error) {

	var mailClient MailClient

	auth := smtp.PlainAuth("", cfg.Mail.MailUsername, cfg.Mail.MailPasswd, cfg.Mail.MailHost)

	pool, err := email.NewPool(cfg.Mail.MailAddress, POOL_SIZE, auth)
	if err != nil {
		return nil, err
	}
	mailClient.mailPool = pool
	return &mailClient, nil
}

var MailPoolProviderSet = wire.NewSet(NewMailClient)

func (mc *MailClient) SendRemindEmails(tos []string, subject, content string) error {
	e := &email.Email{
		To:      tos,
		From:    EMAIL_FROM,
		Subject: subject,
		HTML:    []byte(content),
		Headers: textproto.MIMEHeader{},
	}
	return mc.mailPool.Send(e, 10)
}
