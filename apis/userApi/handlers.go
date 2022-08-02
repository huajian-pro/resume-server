package userApi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"resume-server/database/dao"
	"resume-server/utils"
	"resume-server/utils/response"
)

// 验证码的 redis 服务
var authCodeRedis = dao.NewRedis("authCode")

type RegisterReq struct {
	Username string `json:"username" validate:"required,min=4,max=18"` // 用户名
	Email    string `json:"email" validate:"required,email"`           // 邮箱
	Password string `json:"password" validate:"required,min=6,max=36"` // 密码
	Code     string `json:"code" validate:"required"`                  // 验证码
}

type GetCodeReq struct {
	Email string `json:"email" validate:"required,email"`
}

// 发送注册请求体 -> 判断用户名和邮箱是否存在 -> 校验验证码 -> 密码加密 -> 注册成功
func Register(c *fiber.Ctx) error {
	var req RegisterReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Fail(err.Error()))
	}

	if err := utils.ValidateParams(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.FailWithData("参数错误", err))
	}

	return c.Status(fiber.StatusOK).JSON(response.Ok("注册成功"))
}

// 填写邮箱 -> 点击获取验证码 -> 后端获取邮箱 -> 随机生成验证码存入redis(有效期5分钟) -> 发送验证码到邮箱
func SendCode(c *fiber.Ctx) error {
	defer authCodeRedis.CloseRedis()

	var req GetCodeReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Fail(err.Error()))
	}
	if err := utils.ValidateParams(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.FailWithData("参数错误", err))
	}

	code := utils.GetRandomCode()
	fmt.Println(code)
	// 在redis存储验证码（60秒），后面5分钟
	authCodeRedis.HashSetRedisKey(req.Email, code, 60)
	// 发送邮件
	ok := utils.Email("邮箱验证码", fmt.Sprintf("【化简】验证码%s，用于注册，5分钟内有效。", code)).Send([]string{req.Email})
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.Fail("发送失败"))
	}

	return c.Status(fiber.StatusOK).JSON(response.Ok("发送成功"))
}
