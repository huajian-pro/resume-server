package userApi

import (
	"github.com/gofiber/fiber/v2"
)

func SayHello(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func SayHi(c *fiber.Ctx) error {
	return c.SendString("Hi!")
}
