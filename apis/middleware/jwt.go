package middleware

import (
	"resume-server/conf"
	"resume-server/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 检查白名单
		for _, v := range conf.Cfg.JWT.WhiteList {
			if v == c.Path() {
				return c.Next()
			}
		}

		var code int = 200
		var message string
		headers := c.GetReqHeaders()
		token := headers["Authorization"]
		claims, err := utils.ParseToken(token)
		if token == "" || err != nil {
			code = 400
			message = "用户认证失败"
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = 401
			message = "用户认证信息过期"
		}

		if code != 200 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    code,
				"message": message,
				"success": false,
			})
		}

		return c.Next()
	}
}
