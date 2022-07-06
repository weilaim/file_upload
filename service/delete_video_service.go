package service

import (
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

// DeleteVideoService 删除视频的服务

type DeleteVideoService struct {
}

func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video

	//把要删除的记录查询出来
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}
	//删除视频动作
	err = model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Code: 50000,
			Msg:    "视频删除失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Msg: "删除成功",
	}

}
