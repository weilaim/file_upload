package service

import (

	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

//获取单个小情书
type GetLoveService struct {
	ID  int64 `form:"id" json:"id"`
	Uid int   `form:"uid" json:"uid"`
}

func (service *GetLoveService) Love() serializer.Response {

	if service.ID < 0 {
		return serializer.Response{
			Code: 430,
			Msg:    "本地用户错误",
		}

	}
	var love model.Loves
	if err := model.DB.Preload("Loveimg").Where("id=?", service.ID).Find(&love).Error; err != nil {
		return serializer.Response{
			Code: 5000,
			Msg:    "连接数据库失败",
			Error:  err.Error(),
		}
	}
	love.Uid = service.Uid //传递请求人的id
	//处理小情书被点击的一系列问题
	love.AddView()

	return serializer.BuildLove(serializer.BuildLo(love), 200)
}
