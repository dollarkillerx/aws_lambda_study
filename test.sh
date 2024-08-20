#!/bin/bash

# 构建 Docker 镜像
docker-compose build

# 启动 LocalStack 和 Lambda 容器
docker-compose up -d

## 在 LocalStack 中创建 Lambda 函数
#aws lambda create-function \
#  --function-name hello-lambda \
#  --package-type Image \
#  --code ImageUri=hello-lambda:latest \
#  --role arn:aws:iam::000000000000:role/lambda-role \
#  --endpoint-url=http://localhost:4566

# 测试 Lambda 函数
#aws lambda invoke \
#  --function-name hello-lambda \
#  --payload '{"name": "Leo"}' \
#  --endpoint-url=http://localhost:4566 \
#  output.txt
#
# 查看输出结果
#cat output.txt

curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"name": "Leo"}'
