package main

import (
	"context"                             // 导入用于处理上下文的包
	"fmt"                                 // 导入格式化 I/O 的包
	"github.com/aws/aws-lambda-go/lambda" // 导入 AWS Lambda Go SDK
)

// 定义事件结构，用于接收 Lambda 输入的 JSON 数据
type MyEvent struct {
	Name string `json:"name"` // JSON 键名与结构体字段映射
}

// 定义响应结构，用于返回 Lambda 输出的 JSON 数据
type MyResponse struct {
	Message string `json:"message"` // JSON 键名与结构体字段映射
}

// Lambda 处理函数，接收事件并返回响应
func handler(ctx context.Context, event MyEvent) (MyResponse, error) {
	// 格式化消息并返回
	return MyResponse{Message: fmt.Sprintf("Hello, %s!", event.Name)}, nil
}

// Lambda 函数入口
func main() {
	lambda.Start(handler) // 启动 Lambda 处理程序
}
