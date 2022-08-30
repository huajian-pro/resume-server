package resumeApi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"resume-server/database/model"
)

// FindResumeByUser 查询用户的简历
// GET /v1/resume/find?userid=""
func FindResumeByUser(c *fiber.Ctx) error {
	// 获取请求参数
	userid := c.Query("userid")
	// 查询数据
	var resume model.ResumeData
	resumeList, err := resume.FindAllResumeByBelong(userid)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(resumeList)
}

// SaveResumeData 保存
// POST /v1/resume/save
func SaveResumeData(c *fiber.Ctx) error {
	var ss model.ResumeData
	_ = c.BodyParser(&ss)
	ss.Belong = "62e629c206234829cf6fa130"
	temp, err := ss.InsertResume()
	if err != nil {
		fmt.Println("err:", err)
	}
	return c.JSON(&fiber.Map{"mind": "你发送的POST请求成功了", "temp": temp})
}
