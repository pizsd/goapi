package mail

import (
	"goapi/pkg/config"
	"sync"
)

type From struct {
	Address string
	Name    string
}
type Email struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte
	HTML    []byte
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer // 定义内部使用的Mail，需要使用指针指向的值

func NewMailer() *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: &SMTP{},
		}
	})
	return internalMailer
}

func (m *Mailer) Send(email Email) bool {
	return m.Driver.Send(email, config.GetStringMapString("mail.smtp"))
}
