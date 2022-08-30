package userApi

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"resume-server/database/model"
	"resume-server/utils"
	"resume-server/utils/response"
)

// Login 登录
func Login(c *fiber.Ctx) error {
	req := struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password" validate:"required"`
	}{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Fail(err.Error()))
	}

	if err := utils.ValidateParams(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.FailWithData("参数错误", err))
	}
	// 邮箱和用户名不能同时为空
	if req.Email == "" && req.Username == "" {
		return c.JSON(response.Fail("账号不能为空"))
	}

	var u model.User
	var token string
	u.Password = utils.MD5V([]byte(req.Password))
	if req.Username != "" {
		u.Username = req.Username // 通过唯一用户名找到用户
		if user, err := u.LoginByUsername(); errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(response.Fail("用户名或密码错误"))
		} else {
			token, _ = utils.GenerateToken(user.ID, user.Username, user.Email, user.Phone, user.Role, user.Status)
		}
	}

	if req.Email != "" {
		u.Email = req.Email // 通过唯一用户邮箱找到用户
		if user, err := u.LoginByEmail(); errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(response.Fail("邮箱或密码错误"))
		} else {
			token, _ = utils.GenerateToken(user.ID, user.Username, user.Email, user.Phone, user.Role, user.Status)
		}
	}

	return c.Status(fiber.StatusOK).JSON(response.OkWithData("登录成功", fiber.Map{
		"token": token,
	}))
}
