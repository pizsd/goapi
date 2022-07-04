package mail

import (
	"fmt"
	emailPKG "github.com/jordan-wright/email"
	"github.com/pizsd/goapi/pkg/config"
	"github.com/pizsd/goapi/pkg/helpers"
	"github.com/pizsd/goapi/pkg/logger"
	"net/smtp"
)

type SMTP struct{}

func (st *SMTP) Send(email Email, config map[string]string) bool {
	e := emailPKG.NewEmail()
	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML
	logger.DebugJSON("发送邮件", "邮件详情", e)
	logger.DebugJSON("发送邮件", "config", config)
	err := e.Send(fmt.Sprintf("%v:%v", config["host"], config["port"]), smtp.PlainAuth("", config["username"], config["password"], config["host"]))
	if err != nil {
		logger.ErrorString("发送邮件", "发送出错", err.Error())
		return false
	}
	logger.DebugString("发送邮件", "发送成功", "")
	return true
}

func (st *SMTP) Config() map[string]string {
	stmp := config.GetStringMapString("mail.smtp")
	from := config.GetStringMapString("mail.from")
	return helpers.MapMerge(stmp, from)
}
