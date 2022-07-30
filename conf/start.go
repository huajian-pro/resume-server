package conf

import (
	"flag"
	"fmt"
)

// ---------------------------------------------
// 启动配置
// ---------------------------------------------

// S P T 启动变量
var S = flag.Bool("s", false, "true为正式环境，默认false测试或开发环境")
var P = flag.Bool("p", false, "true为启用多线程，默认false")

// 初始化配置
var config = MyConfig{}

// Cfg 配置信息
var Cfg = config.getMyConfig().EnvTest // 默认测试环境配置

// 初始化配置信息
func init() {
	// 解析命令行参数
	flag.Parse()

	// 设置为发布模式
	if *S == true {
		Cfg = config.getMyConfig().EnvProd // 赋值为生产环境配置
		fmt.Println("各部门注意 -> 🔥生产环境🔥 <- 已启动")
	} else {
		fmt.Println("各部门注意 -> 🌲开发环境🌲 <- 已启动")
	}
}
