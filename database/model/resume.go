package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"resume-server/database/access"
)

// 引入数据库操作接口
var resumeSet = access.ResumeData

// ResumeData 简历数据
type ResumeData struct {
	Belong      string         `bson:"belong" json:"belong"`            // 简历所属用户ID
	TmpID       string         `bson:"tmpID" json:"ID"`                 // 简历模版ID
	Name        string         `bson:"name" json:"NAME"`                // 简历模版名称
	Title       string         `bson:"title" json:"TITLE"`              // 简历模版标题
	Layout      string         `bson:"layout" json:"LAYOUT"`            // 简历模版布局
	Components  []any          `bson:"components" json:"COMPONENTS"`    // 模版组件
	GlobalStyle map[string]any `bson:"globalStyle" json:"GLOBAL_STYLE"` // 模版全局样式
	CreateTime  int64          `bson:"createTime" json:"createTime"`    // 创建时间
	UpdateTime  int64          `bson:"updateTime" json:"updateTime"`    // 更新时间
}

// FindResumeByID 根据ID查询一条数据
func (r *ResumeData) FindResumeByID() (*ResumeData, error) {
	var resume ResumeData
	err := resumeSet.FindOne(
		context.TODO(),
		bson.M{"tmpID": r.TmpID, "belong": r.Belong},
	).Decode(&resume)
	if err != nil {
		return nil, err
	}
	return &resume, nil
}

// FindResumeByName 根据名称查询一条数据
func (r *ResumeData) FindResumeByName(name string) (*ResumeData, error) {
	var resume ResumeData
	err := resumeSet.FindOne(context.TODO(), bson.M{"name": name}).Decode(&resume)
	if err != nil {
		return nil, err
	}
	return &resume, nil
}

// FindAllResumeByBelong 根据所属用户ID查询所有数据
func (r *ResumeData) FindAllResumeByBelong(belong string) ([]ResumeData, error) {
	if belong == "" {
		log.Fatalln("所属用户ID不能为空")
	}
	var resumes []ResumeData
	cur, err := resumeSet.Find(context.TODO(), bson.M{"belong": belong})
	if err != nil {
		return nil, err
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		_ = cur.Close(ctx)
	}(cur, context.TODO())
	for cur.Next(context.TODO()) {
		var resume ResumeData
		err = cur.Decode(&resume)
		if err != nil {
			fmt.Println("err:", err)
		}
		resumes = append(resumes, resume)
	}
	return resumes, nil
}

// InsertResume 插入一条数据
func (r *ResumeData) InsertResume() (*mongo.InsertOneResult, error) {
	return resumeSet.InsertOne(context.TODO(), r)
}

// UpdateResume 更新一条数据
func (r *ResumeData) UpdateResume() (*mongo.UpdateResult, error) {
	return resumeSet.UpdateOne(
		context.TODO(),
		bson.M{"tmpID": r.TmpID, "belong": r.Belong},
		bson.M{"$set": r},
	)
}
