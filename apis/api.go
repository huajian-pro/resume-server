package apis

import (
	"github.com/gofiber/fiber/v2"
	"resume-server/apis/user_api"
)

// Api 设置路由
// 代码便是文档，命名就是路由
func Api(app *fiber.App) {
	v1 := app.Group("v1") // v1 路由组，使用该路由组时，前面需加上 /v1
	// v1.Use()  // 基础的中间件

	// 用户模块
	user := v1.Group("user")
	// user.Use() // 公共路由的中间件
	user.Get("/hello", user_api.SayHello) // 访问：/v1/user/hello
	user.Get("/hi", user_api.SayHi)

	// 简历模块
	resume := v1.Group("resume")
	// resume.Use() // 公共路由的中间件
	resume.Get("/hello", user_api.SayHello) // 访问：/v1/resume/hello
	resume.Get("/hi", user_api.SayHi)
}

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
