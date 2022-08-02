package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"resume-server/apis"
	"resume-server/conf"
)

func main() {
	// 创建引擎
	app := fiber.New(fiber.Config{
		AppName: "Resume Server v0.1", // 应用名称
		Prefork: *conf.P,              // 是否启用多线程，多线程监听同一端口
	})

	// 使用日志
	// app.Use(logger.Default)
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05", // 时间格式
	}))

	// 发送验证码
	// ok := utils.Email("这是标题", "这是内容").Send([]string{"xifive@163.com"})
	// fmt.Println("发送邮件：", ok)

	// 验活
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("👊 Yes, Iam working!")
	})

	// 注册路由组
	apis.Api(app)

	// 启动引擎
	err := app.Listen(":3000")
	if err != nil {
		log.Panic(err)
	}
}
