package apis

import (
	"github.com/gofiber/fiber/v2"
	"resume-server/apis/middle"
	"resume-server/apis/resumeApi"
	"resume-server/apis/userApi"
)

// Api 设置路由
// 代码便是文档，命名就是路由
func Api(app *fiber.App) {
	v1 := app.Group("v1") // v1 路由组，使用该路由组时，前面需加上 /v1

	// 用户模块
	user := v1.Group("user")
	user.Post("/register", userApi.Register)               // 注册 api，访问：/v1/user/register
	user.Post("/authCode", userApi.SendCode)               // 验证码 api，访问：/v1/user/authCode
	user.Post("/login", middle.CheckAuth(), userApi.Login) // 登录api，访问：/v1/user/login

	// 简历模块
	resume := v1.Group("resume", middle.CheckAuth()) // 使用中间件验证用户是否登录
	resume.Get("/find", resumeApi.FindResumeByUser)  // 查询简历，访问：/v1/resume/find
	resume.Post("/save", resumeApi.SaveResumeData)   // 保存简历，访问：/v1/resume/save
}

// todo List
// 用户注册（邮箱+验证+设置密码；发送验证码）
// 用户登录（邮箱+密码）
// 找回密码或重置密码（邮箱+验证+设置密码；异步发送验证码）
// 用户登出（清除登录状态）
// 删除简历
