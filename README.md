# aws_lambda_study

- aws-lambda-rie  
- golang

### 释意

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

### test

``` 
bash test.sh
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