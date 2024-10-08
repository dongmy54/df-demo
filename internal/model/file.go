package model

import "github.com/gogf/gf/v2/net/ghttp"

type FileUploadInput struct {
	File *ghttp.UploadFile `json:"file"`
}

type FileUploadOutput struct {
	FileName string `json:"file_name"`
}
