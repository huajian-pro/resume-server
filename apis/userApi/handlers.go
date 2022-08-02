package userApi

import (
  "github.com/gofiber/fiber/v2"
  "resume-server/cons"
  "resume-server/database/dao"
  "resume-server/utils"
  "resume-server/utils/response"
)

// 发送注册请求体 -> 判断用户名和邮箱是否存在 -> 校验验证码 -> 密码加密 -> 注册成功
func Register(c *fiber.Ctx) error {
  var req struct {
    Username string `json:"username" validate:"required,min=4,max=18"` // 用户名
    Email    string `json:"email" validate:"required,email"`           // 邮箱
    Password string `json:"password" validate:"required,min=6,max=36"` // 密码
    Code     string `json:"code" validate:"required"`                  // 验证码
  }

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
  var req struct {
    Email string `json:"email" validate:"required,email"`
  }

  if err := c.BodyParser(&req); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(response.Fail(err.Error()))
  }
  if err := utils.ValidateParams(&req); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(response.FailWithData("参数错误", err))
  }

  // 验证码的 redis 服务
  var authCodeRedis = dao.NewRedis(req.Email)
  defer authCodeRedis.CloseRedis()

  code := utils.GetRandomCode()
  // 在redis存储验证码 5 分钟
  authCodeRedis.SetRedisKey(code, 60*5)

  //发送邮件
  if !cons.SendAuthCode(req.Email, code) {
    return c.Status(fiber.StatusBadRequest).JSON(response.Fail("发送失败"))
  }

  return c.Status(fiber.StatusOK).JSON(response.Ok("发送成功"))
}
