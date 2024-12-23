# 构建阶段
FROM golang:1.22-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装基础工具
RUN apk add --no-cache git

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/universal_server/main.go

# 运行阶段
FROM alpine

WORKDIR /app

# 创建配置目录
RUN mkdir -p /app/conf

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/main .

# 复制配置文件
COPY conf/dev/conf.yaml /app/conf/

# 暴露应用端口
EXPOSE 9898

# 设置环境变量
ENV CONFIG_PATH=/app/conf/conf.yaml

# 运行应用
ENTRYPOINT ["./main"]
