package service

import (
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

//用户列表服务
type ListUserService struct {
	Limit int `form:"limit"` //多少个用户
	Start int `form:"start"`
}

//List User
func (service *ListUserService) List() serializer.Response {
	users := []model.User{}
	total = 0
	if service.Limit == 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.User{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}
	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&users).Error; err != nil {
		return serializer.Response{
			Code: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildUsers(users), uint(total),200)
}
