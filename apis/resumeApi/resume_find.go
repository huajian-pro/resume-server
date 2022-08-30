package resumeApi

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"resume-server/database/model"
	"resume-server/utils/resp"
)

// FindResumeByUser 查询用户的简历
// GET /v1/resume/find?tmpID=1
func FindResumeByUser(c *fiber.Ctx) error {
	// 获取请求数据
	tmpID := c.Query("tmpID")
	fmt.Println("tmpID:", tmpID)
	// 从中间件中取出用户id
	userID := c.Locals("Userid").(string)
	fmt.Println("userID:", userID)

	resume := model.ResumeData{
		Belong: userID,
		TmpID:  tmpID,
	}
	// 如果模版id为空，则查询所有简历
	if tmpID == "" {
		// 查询所有简历
		if resumeList, err := resume.FindAllResumeByBelong(userID); errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(resp.With(resp.ErrServer, err))
		} else {
			return c.JSON(resp.With(resp.EthIsOK, resumeList))
		}
	} else {
		// 查询指定模版的简历
		if resumeTmp, err := resume.FindResumeByID(); errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(resp.With(resp.ResumeErr, err))
		} else {
			return c.JSON(resp.With(resp.EthIsOK, resumeTmp))
		}
	}
}
