package service


import (
	"fmt"

	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

type CreateLoveService struct {
	Uuid         string   `form:"uuid" json:"uuid"`
	Nick         string   `form:"nick" json:"nick"`
	Sex          int      `form:"sex" json:"sex"`
	Content      string   `form:"content" json:"content"`
	Lve          string   `form:"lve" json:"lve"`
	Professional string   `form:"professional" json:"professional"`
	Img          []string `form:"img" json:"img"`
	Avatar       string   `form:"avatar" json:"avatar"`
}


func (service *CreateLoveService) Create() serializer.Response {
	if service.Professional == "" {
		return serializer.Response{
			Code: 420,
			Msg:    "请填写专业",
		}
	}

	if service.Content == "" {
		return serializer.Response{
			Code: 403,
			Msg:    "请填写内容",
		}
	}



	Loves := model.Loves{
		Uuid:         service.Uuid,
		Nick:         service.Nick,
		Sex:          service.Sex,
		Content:      service.Content,
		Lve:          service.Lve,
		Professional: service.Professional,
		Avatar: service.Avatar,
	}

	err := model.DB.Create(&Loves).Error
	fmt.Println(Loves.ID)
	if err != nil {
		return serializer.Response{
			Code: 5000,
			Msg:    "保存情书失败",
		}
	}

	for _, v := range service.Img {
		//存入url
		LoveImg := model.Loveimg{
			LovesId: int(Loves.ID),
			LoveUrl: v,
		}
		err := model.DB.Create(&LoveImg).Error
		if err != nil {
			return serializer.Response{
				Code: 50001,
				Msg:    "图片链接保存失败",
				Error:  err.Error(),
			}
		}

	}

	return serializer.Response{
		Code: 200,
		Msg:    "创建成功啦",
		Data: map[string]int64{
			"id": int64(Loves.ID),
		},
	}
}
