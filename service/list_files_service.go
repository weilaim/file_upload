package service

import (

	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

type ListFilesService struct {
	Page  int    `form:"page" json:"page"`
	Limit int    `form:"limit" json:"limit"`
	Key   string `form:"key"`
}

//计算File总页数
var fileTotal int64

func (service *ListFilesService) List() serializer.Response {

	if service.Page <= 1 {
		service.Page = 1 //如果起始页小于或等于0
	}
	//初始化 总页数
	fileTotal = 0
	files := []model.Files{}

	if err := model.DB.Limit(service.Limit).Offset((service.Page - 1) * service.Limit).Find(&files).Error; err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "查询文件流错误",
			Error: err.Error(),
		}
	}

	if err := model.DB.Model(model.Files{}).Count(&fileTotal).Error; err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "查询出错了。",
			Error: err.Error(),
		}
	}

	//if total = 0
	if fileTotal == 0 {
		return serializer.Response{
			Msg:  "查询不到记录",
			Code: 203,
		}
	}

	return serializer.BuildListResponse(serializer.BuildFiles(files), uint(fileTotal), 200)
 

}
