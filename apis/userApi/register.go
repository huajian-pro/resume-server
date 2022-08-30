package userApi

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"resume-server/database/dao"
	"resume-server/database/model"
	"resume-server/utils"
	"resume-server/utils/response"
	"time"
)

// Register 注册
func Register(c *fiber.Ctx) error {
	var req struct {
		Username string `json:"username" validate:"required,min=4,max=18"` // 用户名
		Email    string `json:"email" validate:"required,email"`           // 邮箱
		Password string `json:"password" validate:"required,min=6,max=36"` // 密码
		Code     string `json:"code" validate:"required"`                  // 验证码
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Fail(err.Error()))
	}

	if err := utils.ValidateParams(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.FailWithData("参数错误", err))
	}

	var u model.User
	u.Username = req.Username
	u.Email = req.Email
	// 验证码的 redis 服务
	var authCodeRedis = dao.NewRedis(req.Email)
	defer authCodeRedis.CloseRedis()

	if _, err := u.FindUserByUserName(); !errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(response.Fail("用户已存在"))
	}
	if _, err := u.FindUserByEmail(); !errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(response.Fail("邮箱已存在"))
	}

	// 校验验证码
	vcode := authCodeRedis.GetRedisKey()
	if vcode == nil || req.Code != vcode {
		return c.JSON(response.Fail("验证码错误或已过期"))
	}

	// 密码加密
	u.Password = utils.MD5V([]byte(req.Password))

	// 默认普通用户
	u.Role = 0
	u.CreateTime = time.Now().Unix()
	u.UpdateTime = time.Now().Unix()
	u.ID = bson.NewObjectId().Hex()
	_, err := u.CreateUser()
	if err != nil {
		fmt.Println(err)
		return c.JSON(response.Fail("注册失败"))
	}
	// 从 redis 中删除验证码
	authCodeRedis.DelRedisKey()
	token, _ := utils.GenerateToken(u.ID, u.Username, u.Email, u.Phone, u.Role, u.Status)

	return c.Status(fiber.StatusOK).JSON(response.OkWithData("注册成功", fiber.Map{
		"token": token,
	}))
}
