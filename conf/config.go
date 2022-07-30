package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// ---------------------------------------------
// 参数配置
// ---------------------------------------------

// MyConfig 配置文件结构体
type MyConfig struct {
	// 测试环境、开发环境
	EnvTest struct {
		Port  string `yaml:"port"` // 端口配置
		MySQL struct {
			Addr string `yaml:"addr"` // IP地址
			User string `yaml:"user"` // 用户
			Port string `yaml:"port"` // 端口
			Pass string `yaml:"pass"` // 密码
			Base string `yaml:"base"` // 库名
		} `yaml:"mysql"` // mysql配置
		Mongo struct {
			Addr string `yaml:"addr"` // IP地址
			Port string `yaml:"port"` // 端口
			Base string `yaml:"base"` // 库名
		} `yaml:"mongo"` // mongo配置
	} `yaml:"env_test"`
	// 生产环境
	EnvProd struct {
		Port  string `yaml:"port"` // 端口配置
		MySQL struct {
			Addr string `yaml:"addr"` // IP地址
			User string `yaml:"user"` // 用户
			Port string `yaml:"port"` // 端口
			Pass string `yaml:"pass"` // 密码
			Base string `yaml:"base"` // 库名
		} `yaml:"mysql"` // mysql配置
		Mongo struct {
			Addr string `yaml:"addr"` // IP地址
			Port string `yaml:"port"` // 端口
			Base string `yaml:"base"` // 库名
		} `yaml:"mongo"` // mongo配置
	} `yaml:"env_prod"`
}

// 读取配置并绑定结构体
func (m *MyConfig) getMyConfig() *MyConfig {
	// 读取yaml文件到缓存
	yamlFile, err := ioutil.ReadFile("conf/config.yaml")
	if err != nil {
		fmt.Println("没有找到配置文件：", err)
	}
	// yaml文件内容映射到结构体中
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		fmt.Println("绑定配置参数错误：", err.Error())
	}
	return m
}
