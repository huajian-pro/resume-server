package apis

import (
	"github.com/gofiber/fiber/v2"
)

// Api 设置路由
func Api(app *fiber.App) {
	v1 := app.Group("/v1")   // v1 路由组，使用该路由组时，前面需加上 /v1
	v1.Get("/hi", SayHi)     // 使用 /v1/hi 可访问
	app.Get("/hello", SayHi) // 使用 /hello 可访问
}

func SayHi(c *fiber.Ctx) error {
	return c.SendString("Hi!")
}

// 用户登录
// 查询用户的简历列表
// 查询用户的简历详情
