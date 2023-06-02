package mail

import (
	"gin_api_frame/config"
	"net/smtp"

	"github.com/google/wire"
	"github.com/jordan-wright/email"
)

const (
	POOL_SIZE = 4
)

var MailPool *email.Pool

func NewRedisPool(cfg *config.Config) (*email.Pool, error) {

	auth := smtp.PlainAuth("", cfg.Mail.MailUsername, cfg.Mail.MailPasswd, cfg.Mail.MailHost)

	pool, err := email.NewPool(cfg.Mail.MailAddress, POOL_SIZE, auth)
	if err != nil {
		return nil, err
	}
	MailPool = pool
	return pool, nil
}

var MailPoolProviderSet = wire.NewSet(NewRedisPool)

