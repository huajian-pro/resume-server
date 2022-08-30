package userApi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"resume-server/cons"
	"resume-server/database/dao"
	"resume-server/utils"
	"resume-server/utils/response"
)

// SendCode 发送验证码
func SendCode(c *fiber.Ctx) error {
	var req struct {
		Email string `json:"email" validate:"required,email"`
	}
	// 如果获取请求参数失败，则返回错误
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Fail(err.Error()))
	}
	// 如果参数验证失败，则返回错误
	if err := utils.ValidateParams(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.FailWithData("参数错误", err))
	}

	// 验证码的 redis 服务
	var authCodeRedis = dao.NewRedis(req.Email)
	defer authCodeRedis.CloseRedis()

	code := utils.GetRandomCode()
	// 在redis存储验证码 5 分钟
	authCodeRedis.SetRedisKey(code, 60*5)
	fmt.Println("验证码：", code)

	// 发送邮件
	if !cons.SendAuthCode(req.Email, code) {
		return c.Status(fiber.StatusBadRequest).JSON(response.Fail("发送失败"))
	}

	return c.Status(fiber.StatusOK).JSON(response.Ok("发送成功"))
}
