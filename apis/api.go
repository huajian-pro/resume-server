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
// 用户的简历列表
// 用户的简历详情

// 全部模版简历
// 创建简历（使用模版）
// 编辑简历（模块开关及调序、样式配置及数据内容）
// 删除简历
