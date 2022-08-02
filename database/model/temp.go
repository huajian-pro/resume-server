package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"resume-server/database/access"
)

// 引入数据库操作接口
var tempSet = access.TempSet

// Template 模版
// 模版只允许读取，不允许更新
type Template struct {
	ID      string `bson:"_id" json:"TempID"`      // 模版ID
	Name    string `bson:"name" json:"name"`       // 模版名称
	Content string `bson:"content" json:"content"` // 模版内容
}

// FindTempByID 根据ID查询一个模版
func (t *Template) FindTempByID(id string) (*Template, error) {
	var temp Template
	err := tempSet.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&temp)
	if err != nil {
		return nil, err
	}
	return &temp, nil
}

// FindTempByName 根据名称查询一个模版
func (t *Template) FindTempByName(name string) (*Template, error) {
	var temp Template
	err := tempSet.FindOne(context.TODO(), bson.M{"name": name}).Decode(&temp)
	if err != nil {
		return nil, err
	}
	return &temp, nil
}

// InsertTemp 插入一条模版
// 仅初始化时被调用 todo 初始化时插入模版数据
func (t *Template) InsertTemp() (*mongo.InsertOneResult, error) {
	return tempSet.InsertOne(context.TODO(), t)
}
