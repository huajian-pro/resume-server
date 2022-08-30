package middle

import (
	"github.com/gofiber/fiber/v2"
	"resume-server/conf"
	"resume-server/utils/jwt"
	"resume-server/utils/resp"
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

		headers := c.GetReqHeaders()
		token := headers["Authorization"]
		claims, err := jwt.ParseToken(token)
		if token == "" || err != nil {
			return c.JSON(resp.With(resp.AuthFail, nil))
		} else if time.Now().Unix() > claims.ExpiresAt {
			return c.JSON(resp.With(resp.AuthExpired, nil))
		}

		// 传递到下个中间件或处理函数
		c.Locals("Userid", claims.Userid)     // 设置用户id
		c.Locals("Username", claims.Username) // 设置用户名
		c.Locals("Email", claims.Email)       // 设置邮箱
		c.Locals("Role", claims.Role)         // 设置角色
		return c.Next()
	}
}
