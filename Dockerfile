FROM golang:1.18

# 工作目录
WORKDIR /app

# 拷贝文件
COPY . .

# 环境和依赖
RUN go env -w GOPROXY="https://goproxy.cn,direct" \
    && go mod tidy \
    && go build -o /app/resume

# 端口
EXPOSE 3000

# 启动
CMD ["make", "run-prod"]
