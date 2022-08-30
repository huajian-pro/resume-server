package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"resume-server/database/access"
	"time"
)

// 引入数据库操作接口
var resumeSet = access.ResumeData

// ResumeData 模版
// 模版只允许读取，不允许更新
type ResumeData struct {
	Belong      string         `bson:"belong" json:"belong"`            // 简历所属用户ID
	TmpID       string         `bson:"tmpID" json:"ID"`                 // 简历模版ID
	Name        string         `bson:"name" json:"NAME"`                // 简历模版名称
	Title       string         `bson:"title" json:"TITLE"`              // 简历模版标题
	Layout      string         `bson:"layout" json:"LAYOUT"`            // 简历模版布局
	Components  []any          `bson:"components" json:"COMPONENTS"`    // 模版组件
	GlobalStyle map[string]any `bson:"globalStyle" json:"GLOBAL_STYLE"` // 模版全局样式
	CreateTime  time.Time      `bson:"createTime" json:"createTime"`    // 创建时间
	UpdateTime  time.Time      `bson:"updateTime" json:"updateTime"`    // 更新时间
}

// FindResumeByID 根据ID查询一条数据
func (r *ResumeData) FindResumeByID(id string) (*ResumeData, error) {
	err := resumeSet.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// FindResumeByName 根据名称查询一条数据
func (r *ResumeData) FindResumeByName(name string) (*ResumeData, error) {
	var temp ResumeData
	err := resumeSet.FindOne(context.TODO(), bson.M{"name": name}).Decode(&temp)
	if err != nil {
		return nil, err
	}
	return &temp, nil
}

// FindAllResumeByBelong 根据所属用户ID查询所有数据
func (r *ResumeData) FindAllResumeByBelong(belong string) ([]ResumeData, error) {
	var resumes []ResumeData
	cur, err := resumeSet.Find(context.TODO(), bson.M{"belong": belong})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var resume ResumeData
		err := cur.Decode(&resume)
		if err != nil {
			return nil, err
		}
		resumes = append(resumes, resume)
	}
	return resumes, nil
}

// InsertResume 插入一条数据
func (r *ResumeData) InsertResume() (*mongo.InsertOneResult, error) {
	return resumeSet.InsertOne(context.TODO(), r)
}
