package cons

import (
	"fmt"
	"resume-server/utils"
)

// SendAuthCode 发送验证码到邮箱
func SendAuthCode(email, code string) bool {
	return utils.Email().Send("邮箱验证码", fmt.Sprintf("【化简】验证码%s，用于注册，5分钟内有效。", code)).To([]string{email})
}
