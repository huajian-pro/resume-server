package apis

import (
	"github.com/gofiber/fiber/v2"
	"resume-server/conf"
	"resume-server/handlers"
)

// Api 设置路由
func Api(app *fiber.App) {
	baseRoutes := app.Group(conf.Cfg.GlobalPrefix) // v1 路由组，使用该路由组时，前面需加上 /api/v1
	// 基础的中间件
	//baseRoutes.Use()

	publicRoutes := baseRoutes.Group("")
	// 公共路由的中间件
	//publicRoutes.Use()
	{
		// hello
		publicRoutes.Get("/hello", helloHandler.SayHello) // 访问：/api/v1/hello
		publicRoutes.Get("/hi", helloHandler.SayHi)
	}

	privateRoutes := baseRoutes.Group("")
	// 私有路由的中间件
	//privateRoutes.Use()
	{
		privateRoutes.Get("/test", func(ctx *fiber.Ctx) error {
			return ctx.SendString("private")
		})
	}
}

// 在这里注册一下handler，方便读取
var (
	helloHandler = handlers.GroupApp.Hello
)

// todo List
// 用户注册（邮箱+验证+设置密码；异步发送验证码）
// 用户登录（邮箱+密码）

// 找回密码或重置密码（邮箱+验证+设置密码；异步发送验证码）
// 用户登出（清除登录状态）

// 用户的简历列表
// 用户的简历详情
// 全部模版简历（每个模版只能创建一个简历）
// 创建简历（使用模版）
// 编辑简历（模块开关及调序、样式配置及数据内容）
// 删除简历
