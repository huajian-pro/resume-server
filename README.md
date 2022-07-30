# resume-server

## 目录介绍

```text
.
├── README.md            # 说明文件
├── apis                 # API
│   └── api.go           # 统一的api管理入口
│   └── ....go           # 其他api文件
├── conf                 # 配置文件
│   ├── config.go        # 配置文件入口
│   ├── config.yaml      # 配置文件
│   └── start.go         # 启动文件入口
├── cons                 # 常量（状态码、等固定参数）
│   └── cons.go          # 常量文件入口
│   └── ....go           # 其他常量文件，如结构体等
├── database             # 数据库
│   ├── access           # 数据库
│   │   ├── mongo.go     # mongo数据库访问
│   │   ├── mysql.go     # mysql数据库访问
│   │   └── redis.go     # redis数据库访问
│   └── model            # 数据库模型
│       └── ....go       # 数据模型及其操作
├── go.mod               # go模块文件
├── go.sum               # go模块文件
├── http                 # http请求
│   └── self_api.http    # 自定义http请求文件
├── main.go              # 主程序入口文件
└── utils                # 工具类（包含内部或外部）
    └── converter.go     # 转换工具
    └── ...              # 其他工具或文件夹
```

