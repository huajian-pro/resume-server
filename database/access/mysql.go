package access

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"resume-server/conf"
	"time"
)

// DB 通过init直接初始化
var DB *gorm.DB

func init() {
	DB = initMysql()
}

// 定义db的连接信息
var (
	dbUser = conf.Cfg.MySQL.User // 数据库用户名
	dbPass = conf.Cfg.MySQL.Pass // 数据库密码
	dbAddr = conf.Cfg.MySQL.Addr // 数据库地址
	dbPort = conf.Cfg.MySQL.Port // 数据库端口
	dbBase = conf.Cfg.MySQL.Base // 数据库名
)

// 初始化数据库
func initMysql() (db *gorm.DB) {
	// 拼接数据库连接信息
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbAddr, dbPort, dbBase)
	// 初始化db
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer 日志输出的目标，前缀和日志包含的内容
			logger.Config{
				SlowThreshold:             time.Second,   // 慢 SQL 阈值
				LogLevel:                  logger.Silent, // 日志级别
				IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,         // 禁用彩色打印
			}, // 日志配置
		), // Gorm SQL 日志全局模式
		SkipDefaultTransaction: true, // 默认事务，禁用可提升性能
		PrepareStmt:            true, // 开启缓存预编译，提高调用速度
	})
	// 处理错误
	if err != nil {
		fmt.Println("mysql打开失败", err)
		return
	}
	// 一个坑，不设置这个参数，gorm会把表名转义后加个s，导致找不到数据库的表
	SqlDB, _ := db.DB()
	// 设置连接池中最大的闲置连接数
	SqlDB.SetMaxIdleConns(10)
	// 设置数据库的最大连接数量
	SqlDB.SetMaxOpenConns(100)
	// 这是连接的最大可复用时间
	SqlDB.SetConnMaxLifetime(10 * time.Second)
	return db
}
