# 创建
FROM alpine:3.6

# 工作目录
WORKDIR /app

# 拷贝文件
COPY resume-lin .

# 导出端口
EXPOSE 3000

# 执行命令
RUN chmod +x ./resume-lin

# 启动
CMD ["./resume"]
