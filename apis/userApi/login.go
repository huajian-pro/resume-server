package userApi

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"resume-server/database/model"
	"resume-server/utils"
	"resume-server/utils/jwt"
	"resume-server/utils/resp"
)

// Login 登录
func Login(c *fiber.Ctx) error {
	req := struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password" validate:"required"`
	}{}
	// 获取并绑定请求参数
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(resp.With(resp.ErrServer, err))
	}
	if err := utils.ValidateParams(&req); err != nil {
		return c.JSON(resp.With(resp.ErrParam, err))
	}
	// 邮箱和用户名不能同时为空
	if req.Email == "" && req.Username == "" {
		return c.JSON(resp.With(resp.LoginFail, nil))
	}

	var u model.User
	var token string
	u.Password = utils.MD5V([]byte(req.Password))
	if req.Username != "" {
		u.Username = req.Username // 通过唯一用户名找到用户
		if user, err := u.LoginByUsername(); errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(resp.With(resp.LoginUserNameErr, nil))
		} else {
			token, _ = jwt.GenerateToken(user.ID, user.Username, user.Email, user.Phone, user.Role, user.Status)
		}
	}

	if req.Email != "" {
		u.Email = req.Email // 通过唯一用户邮箱找到用户
		if user, err := u.LoginByEmail(); errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(resp.With(resp.LoginUserEmailErr, nil))
		} else {
			token, _ = jwt.GenerateToken(user.ID, user.Username, user.Email, user.Phone, user.Role, user.Status)
		}
	}

	return c.JSON(resp.With(resp.EthIsOK, fiber.Map{
		"token": token,
	}))
}
