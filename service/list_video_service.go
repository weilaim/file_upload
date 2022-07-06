package service

import (
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

// ListVideoService 视频详情列表服务
type ListVideoService struct {
	Limit int `form:"limit"` //多少条视频
	Start int `form:"start"`
}

var total int64

// List 视频列表
func (service *ListVideoService) List() serializer.Response {
	videos := []model.Video{}
	total = 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Video{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&videos).Error; err != nil {
		return serializer.Response{
			Code: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total),200)
}
