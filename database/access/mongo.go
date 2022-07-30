package access

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"resume-server/conf"
)

// SetName 是 mongo 集合句柄（支持多个句柄）
var SetName *mongo.Collection // 句柄

// mongo 的连接信息
var (
	Addr = conf.Cfg.Mongo.Addr // 地址
	Port = conf.Cfg.Mongo.Port // 端口
	Base = conf.Cfg.Mongo.Base // 库名
)

// 初始化mongo链接，集合句柄
func init() {
	mgoDB := connectToDB(fmt.Sprintf("mongodb://%v:%v", Addr, Port), Base) // 第二个参数是库名
	// mongo 集合句柄
	SetName = mgoDB.Collection("") // 给句柄赋值，name为集合名
	// ...
}

// 链接mongo连接池
func connectToDB(uri, name string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(50)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return client.Database(name)
}
