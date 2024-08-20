# aws_lambda_study

- aws-lambda-rie  
- localstack 
- golang

### 释意

- aws-lambda-rie
> aws-lambda-rie 是 Amazon Web Services (AWS) 提供的一个工具，称为 "AWS Lambda Runtime Interface Emulator"。
> 它是一个用于本地模拟 AWS Lambda 运行时的轻量级工具，主要用于在 Docker 容器中测试 AWS Lambda 函数，而无需在真正的 AWS Lambda 环境中部署代码。

- localstack 
> LocalStack 是一个用于本地开发的工具，它可以模拟AWS的云服务 
> LocalStack 支持包括S3、Lambda、DynamoDB、SNS、SQS等多种常用的AWS服务。

### demo

``` 
├── LICENSE
├── README.md
├── demo1
│   └── cmd
│       └── lambda
│           └── hello
│               ├── Dockerfile
│               ├── docker-compose.yml
│               ├── main.go
│               └── test.sh
├── go.mod
└── go.sum
```

通过 LocalStack 模拟 AWS 环境，并使用 aws-lambda-rie 来运行和测试这个函数。

``` 
mkdir -p cmd/lambda/hello
```

构建
``` 
docker-compose build
```

启动 LocalStack 和 Lambda 容器：
``` 
docker-compose up -d
```

部署 Lambda 函数到 LocalStack：
使用 aws CLI 将 Lambda 函数部署到 LocalStack：

``` 
# 在 LocalStack 中创建 Lambda 函数
aws lambda create-function \
  --function-name hello-lambda \
  --package-type Image \
  --code ImageUri=hello-lambda:latest \            // 使用镜像部署
  --role arn:aws:iam::000000000000:role/lambda-role \
  --endpoint-url=http://localhost:4566
```
  

测试 Lambda 函数
```
aws lambda invoke \
  --function-name hello-lambda \
  --payload '{"name": "Leo"}' \
  --endpoint-url=http://localhost:4566 \
  output.txt

cat output.txt
``` 


### 依赖

``` 
macos: 
brew install aws/tap/aws-sam-cli

ubuntu:
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
aws --version

wget https://github.com/aws/aws-sam-cli/releases/latest/download/aws-sam-cli-linux-x86_64.zip
unzip aws-sam-cli-linux-x86_64.zip -d sam-installation
sudo ./sam-installation/install
sam --version
```