package service

import (
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

type DeleteFileService struct {

}

func (service *DeleteFileService)Delete(id string) serializer.Response  {
	//要把删除的记录查询出来
	var file model.Files
	err := model.DB.First(&file,id).Error
	if err != nil{
		return serializer.Response{
			Code:404,
			Msg: "文件流不存在",
			Error: err.Error(),
		}
	}

	// 删除视频的动作
	err = model.DB.Delete(&file).Error
	if err != nil {
		return serializer.Response{
			Code: 501,
			Msg: "删除文件流失败啦",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Msg: "删除成功",
	}
}

