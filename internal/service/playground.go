// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IPlayground interface {
		// 这是我自己定义的一个练写方法
		Do()
	}
)

var (
	localPlayground IPlayground
)

func Playground() IPlayground {
	if localPlayground == nil {
		panic("implement not found for interface IPlayground, forgot register?")
	}
	return localPlayground
}

func RegisterPlayground(i IPlayground) {
	localPlayground = i
}
