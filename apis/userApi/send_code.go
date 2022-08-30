package userApi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"resume-server/cons"
	"resume-server/database/dao"
	"resume-server/utils"
	"resume-server/utils/resp"
)

// SendCode 发送验证码
func SendCode(c *fiber.Ctx) error {
	var req struct {
		Email string `json:"email" validate:"required,email"`
	}
	// 获取并绑定请求参数
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(resp.With(resp.ErrParam, err))
	}
	if err := utils.ValidateParams(&req); err != nil {
		return c.JSON(resp.With(resp.ErrParam, err))
	}

	// 验证码的 redis 服务
	var authCodeRedis = dao.NewRedis(req.Email)
	defer authCodeRedis.CloseRedis()
	code := utils.GetRandomCode()         // 随机生成验证码
	authCodeRedis.SetRedisKey(code, 60*5) // 在redis存储验证码 5 分钟
	fmt.Println("验证码：", code)

	// 发送邮件
	if !cons.SendAuthCode(req.Email, code) {
		return c.JSON(resp.With(resp.AuthSendOTPErr, nil))
	}
	return c.JSON(resp.With(resp.EthIsOK, nil))
}
