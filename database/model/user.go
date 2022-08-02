package model

import (
	"context"
	"resume-server/database/access"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 引入数据库操作接口
var userSet = access.UserSet

// User 用户
type User struct {
	ID         string `bson:"_id"`                          // 用户ID
	Avatar     string `bson:"avatar" json:"avatar"`         // 用户头像
	Username   string `bson:"username" json:"username"`     // 用户名
	Nickname   string `bson:"nickname" json:"nickname"`     // 用户昵称
	Password   string `bson:"password" json:"password"`     // 用户密码, md5加密后的
	Email      string `bson:"email" json:"email"`           // 用户邮箱
	Phone      string `bson:"phone" json:"phone"`           // 用户手机号
	Status     int    `bson:"status" json:"status"`         // 用户状态: 0-启用 1-禁用
	Role       int    `bson:"role" json:"role"`             // 用户角色：0-普通用户，1-VIP，9-管理员
	CreateTime int64  `bson:"createTime" json:"createTime"` // 创建时间
	UpdateTime int64  `bson:"updateTime" json:"updateTime"` // 更新时间
}

// CreateUser 创建一个用户
func (u *User) CreateUser() (*mongo.InsertOneResult, error) {
	return userSet.InsertOne(context.TODO(), u)
}

// FindUserByEmail 用邮箱查询一个用户
func (u *User) FindUserByEmail() (*User, error) {
	var user User
	err := userSet.FindOne(context.TODO(), bson.M{"email": u.Email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserByUsername 用username查询一个用户
func (u *User) FindUserByUsername() (*User, error) {
	var user User
	err := userSet.FindOne(context.TODO(), bson.M{"username": u.Username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// LoginByUsername 使用用户名登录
func (u *User) LoginByUsername() (*User, error) {
	var user User
	err := userSet.FindOne(context.TODO(), bson.M{"username": u.Username, "password": u.Password}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// LoginByEmail 使用邮箱登录
func (u *User) LoginByEmail() (*User, error) {
	var user User
	err := userSet.FindOne(context.TODO(), bson.M{"email": u.Email, "password": u.Password}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserByName 用姓名查询一个用户
func (u *User) FindUserByName(name string) (*User, error) {
	var user User
	err := userSet.FindOne(context.TODO(), bson.M{"nickname": name}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
