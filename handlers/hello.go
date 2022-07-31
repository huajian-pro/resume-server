package handlers

import (
	"github.com/gofiber/fiber/v2"
	"resume-server/database/model"
)

type Hello struct {
}

func (h *Hello) SayHello(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func (h *Hello) SayHi(c *fiber.Ctx) error {
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
