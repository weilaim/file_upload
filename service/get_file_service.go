package service

import (
	"fmt"
	"os"

	"github.com/weilaim/blog-api/middleware"
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
	"github.com/weilaim/blog-api/util/errmsg"
)

type GetFileService struct {
	ID    int64  `form:"id" json:"id"`
	Token string `form:"token" json:"token"`
}

func (service *GetFileService) checkFileAuth() int {
	if service.Token == "" {
		return 407
	}
	fmt.Println("aaaa",service.Token)
	_, tokenkCode := middleware.CheckToken(service.Token)
	if tokenkCode == errmsg.ERROR {
		return 406
	}
	return 200
}


//只能有自定义key或者token来访问接口 key:dsflafjlasjf21123qwerety
func (service *GetFileService) SingoFile() serializer.Response {
	AutuKey := os.Getenv("FILE_ACC_KEY")
	if service.Token != AutuKey {
		auth := service.checkFileAuth()
		if auth == 406{
			return serializer.Response{
				Code: 406,
				Msg:  "token不正确",
			}
		}
		if auth == 407{
			return serializer.Response{
				Code: 407,
				Msg:  "token为空",
			}
		}
		 
	}
	if service.ID < 0 {
		return serializer.Response{
			Code: 508,
			Msg:  "查询ID出错",
		}
	}

	var file model.Files
	var findCount int64
	if err := model.DB.Where("id = ?", service.ID).Find(&file).Count(&findCount).Error; err != nil {
		return serializer.Response{
			Code:  503,
			Msg:   "查询错了捏",
			Error: err.Error(),
		}
	}
	if findCount <= 0 {
		return serializer.Response{
			Code: 504,
			Msg:  "查询出错，该内容不存在！",
		}
	}

	return serializer.BuildFileResponse(serializer.BuildFile(file))
}
