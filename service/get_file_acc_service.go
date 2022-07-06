package service

import (
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

type GetFileAccService struct{}

func (service *GetFileAccService) FileAcc(id string) serializer.Response {
	var filecount int64
	var facc model.FilesAcc
	if err := model.DB.Where("id = ?", id).Find(&facc).Count(&filecount).Error; err != nil {
		return serializer.Response{
			Code:  503,
			Msg:   "查询错了捏",
			Error: err.Error(),
		}
	}
	if filecount <= 0 {
		return serializer.Response{
			Code: 503,
			Msg:  "查询失败！该内容不存在",
		}
	}

	return serializer.BuildFileResponse(serializer.BuildFileAcc(facc))
}
