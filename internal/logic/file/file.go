package file

import (
	"context"
	"fmt"
	"gf-demo/internal/model"
	"gf-demo/internal/service"

	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
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
	file := in.File
	file.Filename = fmt.Sprintf("Myfile_%s.txt", gtime.Datetime())
	filename, err := file.Save(gfile.TempDir())
	if err != nil {
		return out, err
	}
	return model.FileUploadOutput{FileName: filename}, err
}
