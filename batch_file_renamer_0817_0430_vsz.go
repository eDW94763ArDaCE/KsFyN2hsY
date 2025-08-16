// 代码生成时间: 2025-08-17 04:30:23
package main

import (
    "context"
    "flag"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
)

// RenameRequest 是批量重命名请求的gRPC消息结构
type RenameRequest struct {
    SourceFiles []string `json:"source_files"`
    NewName    string   `json:"new_name"`
}

// RenameResponse 是批量重命名响应的gRPC消息结构
type RenameResponse struct {
    Result string `json:"result"`
}

// FileRenamerServer 是gRPC服务的接口定义
type FileRenamerServer struct{}

// RenameFiles 实现批量文件重命名功能
func (s *FileRenamerServer) RenameFiles(ctx context.Context, req *RenameRequest) (*RenameResponse, error) {
    // 检查请求参数
    if req == nil || len(req.SourceFiles) == 0 || req.NewName == "" {
        return nil, errInvalidRequest
