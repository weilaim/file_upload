package service

import (

	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

type CreateFileService struct {
	UserId  uint   `form:"user_id" json:"user_id"`
	AccName string `form:"acc_name" json:"acc_name" binding:"required"`
	Title   string `form:"title" json:"title" binding:"required,min=3,max=20"`
	FileUrl string `form:"file_url" json:"file_url"`
	Content string `form:"content" json:"content"`
}

// ValidateTitle
func (service *CreateFileService) ValidateTitle() bool {
	var CheckTitle int64
	model.DB.Model(model.Files{}).Where("title = ?", service.Title).Count(&CheckTitle)
	if CheckTitle > 0 {
		return false
	} else {
		return true
	}

}
func (service *CreateFileService) Files() serializer.Response {
	titleTemp := service.ValidateTitle()
	if !titleTemp {
		return serializer.Response{
			Code: 403,
			Msg:  "该文件名已存在",
		}
	}
	Files := model.Files{
		UserId:  service.UserId,
		Title:   service.Title,
		AccName: service.AccName,
		FileUrl: service.FileUrl,
		Content: service.Content,
	}

	err := model.DB.Create(&Files).Error
	if err != nil {
		return serializer.Response{
			Code: 501,
			Msg:  "文件内容创建失败了~",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "创建成功！",
		Data: serializer.BuildFile(Files),
	}
}
