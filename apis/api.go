package apis

import (
	"github.com/gofiber/fiber/v2"
	"resume-server/database/model"
)

// Api 设置路由
func Api(app *fiber.App) {
	v1 := app.Group("/v1")   // v1 路由组，使用该路由组时，前面需加上 /v1
	v1.Get("/hi", SayHi)     // 使用 /v1/hi 可访问
	app.Get("/hello", SayHi) // 使用 /hello 可访问
}

func SayHi(c *fiber.Ctx) error {
	var u = model.User{}
	u.NickName = "admin"
	u.Password = "123456"
	u.Email = "admin@163.com"
	result, err := u.CreateUser()
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString("Hi!" + result.InsertedID.(string))
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
