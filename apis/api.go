package apis

import (
	"github.com/gofiber/fiber/v2"
	"resume-server/apis/resumeApi"
	"resume-server/apis/userApi"
)

// Api 设置路由
// 代码便是文档，命名就是路由
func Api(app *fiber.App) {
	v1 := app.Group("v1") // v1 路由组，使用该路由组时，前面需加上 /v1
	// v1.Use()  // 基础的中间件

	// 用户模块
	user := v1.Group("user")
	// user.Use() // 公共路由的中间件
	user.Get("/hello", userApi.SayHello) // 访问：/v1/user/hello
	user.Get("/hi", userApi.SayHi)

	// 简历模块
	resume := v1.Group("resume")
	// resume.Use() // 公共路由的中间件
	resume.Get("/find", resumeApi.FindResumeByUser) // 全部模版，访问：/v1/resume/hello
	resume.Post("/save", resumeApi.SaveResumeData)
}

// todo List
// 用户注册（邮箱+验证+设置密码；异步发送验证码）
// 用户登录（邮箱+密码）
// 找回密码或重置密码（邮箱+验证+设置密码；异步发送验证码）
// 用户登出（清除登录状态）

// 全部模版（每个模版只能创建一个简历）
// 使用模版（如果有历史数据则返回之前保存的信息，如果没有历史数据则使用模版数据）
// 保存简历（模块开关及调序、样式配置及数据内容）
// 删除简历
