package resumeApi

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"resume-server/database/model"
	"resume-server/utils/resp"
)

// SaveResumeData 保存
// POST /v1/resume/save
func SaveResumeData(c *fiber.Ctx) error {
	// 获取请求数据
	req := model.ResumeData{}
	req.Belong = c.Locals("Userid").(string)
	// 获取并绑定请求参数
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(resp.With(resp.ErrServer, err))
	}

	// 判断当前用户是否已经存在该模版的简历，如果存在则更新，不存在则新增
	if _, err := req.FindResumeByID(); errors.Is(err, mongo.ErrNoDocuments) {
		// 新增简历
		insetID, insetErr := req.InsertResume()
		if insetErr != nil {
			return c.JSON(resp.With(resp.EthIsOK, insetErr))
		}
		return c.JSON(resp.With(resp.EthIsOK, insetID))
	} else {
		// 更新简历
		updateTmp, updateErr := req.UpdateResume()
		if updateErr != nil {
			return c.JSON(resp.With(resp.EthIsOK, updateErr))
		}
		return c.JSON(resp.With(resp.EthIsOK, updateTmp.ModifiedCount))
	}
}
