package service

import (

	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

//UserInfoService
type UserInfoService struct {
	Id string  `form:"id" json:"id"`
}

func (service *UserInfoService) Get() serializer.Response {
	var user model.User
	if err := model.DB.Where("id = ?", service.Id).First(&user).Error; err != nil {
		return serializer.Response{
			Code:402 ,
			Msg: "未找到该用户",
			Error: err.Error(),
		}
	}
	return serializer.BuildUserResponse(user)
}
