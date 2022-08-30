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
	"resume-server/utils/jwt"
	"resume-server/utils/resp"
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
	// 获取并绑定请求数据出错
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(resp.With(resp.ErrParam, err))
	}
	if err := utils.ValidateParams(&req); err != nil {
		return c.JSON(resp.With(resp.ErrParam, err))
	}

	var u model.User
	u.Username = req.Username
	u.Email = req.Email
	// 验证码的 redis 服务
	var authCodeRedis = dao.NewRedis(req.Email)
	defer authCodeRedis.CloseRedis()
	// 校验用户名
	if _, err := u.FindUserByUserName(); !errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(resp.With(resp.RegisterUserNameErr, nil))
	}
	// 校验邮箱
	if _, err := u.FindUserByEmail(); !errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(resp.With(resp.RegisterUserEmailErr, nil))
	}
	// 校验验证码
	vCode := authCodeRedis.GetRedisKey()
	if vCode == nil || req.Code != vCode {
		return c.JSON(resp.With(resp.RegisterOTPExpired, nil))
	}

	// 创建默认普通用户
	u.Role = 0                                    // 设置用户权限
	u.Password = utils.MD5V([]byte(req.Password)) // 密码加密
	u.CreateTime = time.Now().Unix()              // 创建时间戳
	u.UpdateTime = time.Now().Unix()              // 更新时间戳
	u.ID = bson.NewObjectId().Hex()               // 创建用户id
	_, err := u.CreateUser()                      // 创建用户
	if err != nil {
		fmt.Println(err)
		return c.JSON(resp.With(resp.ErrServer, err))
	}
	// 从 redis 中删除验证码
	authCodeRedis.DelRedisKey()
	token, _ := jwt.GenerateToken(u.ID, u.Username, u.Email, u.Phone, u.Role, u.Status)

	return c.JSON(resp.With(resp.EthIsOK, fiber.Map{
		"token": token,
	}))
}
