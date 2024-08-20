# 使用 Go 1.17.5 作为构建环境
FROM golang:1.17.5-buster as build

# 定义构建参数，用于指定 Lambda 函数名称
ARG lambda_name=hello

# 禁用 CGO，使生成的二进制文件为静态链接
ENV CGO_ENABLED=0

# 下载并安装 AWS Lambda Runtime Interface Emulator (RIE)
ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/local/bin/aws-lambda-rie
RUN chmod 755 /usr/local/bin/aws-lambda-rie  # 赋予执行权限

# 设置工作目录
WORKDIR /app

# 复制 Go 模块文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制项目的所有源文件
COPY . .

# 进入指定的 Lambda 函数目录并构建二进制文件
RUN cd ./cmd/lambda/${lambda_name} && go build -v -o /main

# 设置容器入口点，使用 RIE 运行 Lambda 函数
ENTRYPOINT ["/usr/local/bin/aws-lambda-rie", "/main"]
