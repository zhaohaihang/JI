package mail

import (
	"ji/config"
	"net/smtp"

	"github.com/google/wire"
	"github.com/jordan-wright/email"
)

var MailPool *email.Pool

func NewRedisPool(cfg *config.Config) (*email.Pool, error) {

	auth := smtp.PlainAuth("", "1932859223@qq.com", "mddebbjnnqipbdjg", "smtp.qq.com")

	pool, err := email.NewPool("smtp.qq.com:25", 4, auth)
	if err != nil {
		return nil, err
	}
	MailPool = pool
	return pool, nil
}

var MailPoolProviderSet = wire.NewSet(NewRedisPool)

//mddebbjnnqipbdjg
