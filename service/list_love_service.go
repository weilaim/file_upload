package service

import (
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

type ListLoveService struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

//计算总页数
var lovetotal int64

func (service *ListLoveService) List() serializer.Response {

	if service.Page <= 0 {
		service.Page = 1
	}
	//total 初始化
	lovetotal = 0
	loves := []model.Loves{}

	if err := model.DB.Preload("Loveimg").Limit(service.Limit).Offset((service.Page - 1) * service.Limit).Order("updated_at desc").Find(&loves).Error; err != nil {
		return serializer.Response{
			Code: 5000,
			Msg:    "连接数据库失败",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Model(model.Loves{}).Count(&lovetotal).Error; err != nil {
		return serializer.Response{
			Code: 5000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	//计算总页数

	Page := lovetotal / int64(service.Limit)

	if total%int64(service.Limit) != 0 {
		Page++
	}

	return serializer.BuildListResponse(serializer.BuildLoves(loves), uint(lovetotal), 200)

}
