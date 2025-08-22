// 代码生成时间: 2025-08-23 04:39:23
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"
    "gopkg.in/go-playground/validator.v10"
)

// FormValidationService 定义表单验证服务
type FormValidationService struct {
    // 这里可以添加服务所需要的字段
}

// ValidateFormRequest 定义表单验证请求结构体
type ValidateFormRequest struct {
    Username string `valid:"required,alphanum,min=3,max=30"`
    Email    string `valid:"required,email"`
    Age      int    `valid:"required,min=18"`
}

// ValidateFormResponse 定义表单验证响应结构体
type ValidateFormResponse struct {
    IsValid bool   `json:"isValid"`
    Errors  string `json: