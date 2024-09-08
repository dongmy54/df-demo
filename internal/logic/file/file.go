package file

import (
	"context"
	"fmt"
	"gf-demo/internal/model"
	"gf-demo/internal/service"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/v2/frame/g"
)

type sFile struct{}

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

// Upload 创建内容
func (s *sFile) Upload(ctx context.Context, in model.FileUploadInput) (out model.FileUploadOutput, err error) {
	// file := r.GetUploadFile("TestFile")
	//   if file == nil {
	//       r.Response.Write("empty file")
	//       return
	//   }

	file := in.File
	file.Filename = fmt.Sprintf("Myfile_%s.txt", gtime.Datetime())
	file_dir, _ := g.Cfg().Get(ctx, "file.upload_dir")
	filename, err := file.Save(file_dir.String())
	if err != nil {
		return out, err
	}
	return model.FileUploadOutput{FileName: filename}, err
}
