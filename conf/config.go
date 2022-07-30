package conf

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
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
	var yamlFile []byte
	if *S {
		// 读取生产环境的 yaml 配置文件到缓存
		yamlFile, _ = ioutil.ReadFile("conf/prod.yaml")
	} else {
		// 读取开发环境的 yaml 配置文件到缓存
		yamlFile, _ = ioutil.ReadFile("conf/config.yaml")
	}

	// 如果读取失败，则报错并退出
	if yamlFile == nil {
		log.Fatalln("没有找到配置文件，或者配置文件为空，请检查！")
	}

	// yaml文件内容映射到结构体中，失败则报错并退出
	err := yaml.Unmarshal(yamlFile, m)
	if err != nil {
		log.Fatalln("绑定配置参数错误：", err.Error())
	}
	return m
}
