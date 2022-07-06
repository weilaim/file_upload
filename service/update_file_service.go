package service

import (
	"fmt"

	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

type UpadteFileService struct {
	UserId  uint   `form:"user_id" json:"user_id"`
	Title   string `form:"title" json:"title" `
	FileUrl string `form:"file_url" json:"file_url"`
	Content string `form:"content" json:"content"`
}

func (service *UpadteFileService) Update(id string) serializer.Response {
	var file model.Files
	err := model.DB.First(&file, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "File不存在",
			Error: err.Error(),
		}

	}
	fmt.Println(file.Title)
	fmt.Println(service.Title)
	file.Title = service.Title
	file.FileUrl = service.FileUrl
	file.Content = service.Content
	err = model.DB.Save(&file).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "保存文件流失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Msg: "修改成功",
		Code: 200,
		Data: serializer.BuildFile(file),
	}
}
