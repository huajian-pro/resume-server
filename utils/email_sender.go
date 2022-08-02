package utils

import (
  "fmt"
  "github.com/jordan-wright/email"
  "net/smtp"
  "resume-server/conf"
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
func newEmailSender(from string) *emailSender {
  return &emailSender{
    Mail: email.NewEmail(), // 创建一个邮件对象
    From: from,             // 发件人
  }
}

var once sync.Once             // 一次性锁，防止多个 goroutine 同时调用
var sender *emailSender        // 单例句柄
var emailConf = conf.Cfg.Email // email 配置

// Email 获取单例句柄
func Email() *emailSender {
  once.Do(func() {
    sender = newEmailSender("化简 <" + emailConf.User + ">")
  })
  return sender
}

// Send 发送
func (sender *emailSender) Send(subject, sendBody string) *emailSender {
  sender.Subject = subject   // 主题
  sender.SendBody = sendBody // 文本内容
  return sender
}

// To 发送邮件
func (sender *emailSender) To(toSomebody []string) (ok bool) {
  sender.Mail.From = sender.From             // 发件人
  sender.Mail.To = toSomebody                // 收件人
  sender.Mail.Subject = sender.Subject       // 主题
  sender.Mail.Text = []byte(sender.SendBody) // 文本内容
  err := sender.Mail.Send(
    fmt.Sprintf("%s:%s", emailConf.Host, emailConf.Port), // 邮件服务器地址
    smtp.PlainAuth("", emailConf.User, emailConf.Pass, emailConf.Host),
  )
  if err != nil {
    return false // 发送失败
  }
  return true // 发送成功
}
