package middle

import (
	"github.com/gofiber/fiber/v2"
	"resume-server/conf"
	"resume-server/utils"
	"time"
)

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 检查白名单
		for _, v := range conf.Cfg.JWT.WhiteList {
			if v == c.Path() {
				return c.Next()
			}
		}

		code := 200
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

		// 传递到下个中间件或处理函数
		c.Locals("Userid", claims.Userid)     // 设置用户id
		c.Locals("Username", claims.Username) // 设置用户名
		c.Locals("Email", claims.Email)       // 设置邮箱
		c.Locals("Role", claims.Role)         // 设置角色
		return c.Next()
	}
}
