package controller

import (
	"context"

	"gf-demo/api/backend"
	"gf-demo/internal/model"
	"gf-demo/internal/service"
)

// File 内容管理
var File = cFile{}

type cFile struct{}

// 方法名区分后端和前端
func (a *cFile) BackendUpload(ctx context.Context, req *backend.FileUploadReq) (res *backend.FileUploadRes, err error) {
	out, err := service.File().Upload(ctx, model.FileUploadInput{
		File: req.File,
	})
	if err != nil {
		return nil, err
	}
	return &backend.FileUploadRes{
		FileName: out.FileName,
	}, nil
}
