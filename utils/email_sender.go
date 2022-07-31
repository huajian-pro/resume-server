package utils

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"sync"
)

// 邮件发送者数据结构
type emailSender struct {
	Mail     *email.Email // 邮件对象
	From     string       // 发件人
	Subject  string       // 主题
	SendBody string       // 文本内容
}

// 新建一个邮件发送者
func newEmailSender(from, subject, sendBody string) *emailSender {
	return &emailSender{
		Mail:     email.NewEmail(), // 创建一个邮件对象
		From:     from,             // 发件人
		Subject:  subject,          // 主题
		SendBody: sendBody,         // 文本内容
	}
}

var once sync.Once
var sender *emailSender

// Email 获取单例句柄
func Email(subject, sendBody string) *emailSender {
	once.Do(func() {
		sender = newEmailSender("如果我是dj <xxx@126.com>", subject, sendBody)
	})
	return sender
}

// Send 发送邮件
func (sender *emailSender) Send(toSomebody []string) (ok bool) {
	sender.Mail.From = sender.From             // 发件人
	sender.Mail.To = toSomebody                // 收件人
	sender.Mail.Subject = sender.Subject       // 主题
	sender.Mail.Text = []byte(sender.SendBody) // 文本内容
	err := sender.Mail.Send(
		"smtp.126.com:25",
		smtp.PlainAuth("", "xxx@126.com", "yyy", "smtp.126.com"),
	)
	if err != nil {
		return false // 发送失败
	}
	return true // 发送成功
}

// 发送注册验证码邮件
// 发送重置验证码邮件
