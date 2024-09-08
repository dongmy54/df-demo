package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/file/upload" method:"post" mime:"multipart/form-data" summary:"上传文件" tags:"工具"`
	File   *ghttp.UploadFile `json:"file" type:"file" v:"required#请选择上传文件" dc:"选择上传文件"`
}

type FileUploadRes struct {
	FileName string `json:"file_name"`
}
