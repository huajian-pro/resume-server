package response

import "github.com/gofiber/fiber/v2"

func Ok(message string) fiber.Map {
  return fiber.Map{
    "code":    200,
    "success": true,
    "message": message,
  }
}

func OkWithData(message string, data any) fiber.Map {
  return fiber.Map{
    "code":    200,
    "success": true,
    "message": message,
    "data":    data,
  }
}

func Fail(message string) fiber.Map {
  return fiber.Map{
    "code":    400,
    "success": false,
    "message": message,
  }
}

func FailWithData(message string, data any) fiber.Map {
  return fiber.Map{
    "code":    400,
    "success": false,
    "message": message,
    "data":    data,
  }
}
