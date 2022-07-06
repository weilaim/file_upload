package service

import (
	"fmt"
	"os"

	"github.com/medivhzhan/weapp/v2"
	"github.com/weilaim/blog-api/middleware"
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

type GetOpenIdService struct {
	Code        string `form:"code" json:"code"`
	Nick        string `form:"nick" json:"nick"`
	AvatarUrl   string `form:"avatarurl" json:"avatarurl"`
	Sex         int32  `form:"sex" json:"sex"`
	UnionId     string `form:"unionid" json:"unionid"`
	Openid      string `form:"openid" json:"openid"`
	Session_key string `form:"session_key" json:"session_key"`
}

//ValidUser 验证用户是否存在 100 已经存在用户， 200 则反之
func (service *GetOpenIdService) ValidUser(id string) int {
	var user int64
	model.DB.Model(&model.Wxuser{}).Where("openid = ?", id).Count(&user)
	if user > 0 {

		return 100
	} else {
		return 200
	}
}

func (service *GetOpenIdService) Getid() serializer.Response {
	resp, err := weapp.Login(os.Getenv("WX_APPID"), os.Getenv("WX_SECRET"), service.Code)
	if err != nil {
		return serializer.Response{
			Code: 5004,
			Msg:    "未知错误",
			Error:  err.Error(),
		}
	}

	// 得到了openid 设置token
	token, code := middleware.SetToken(service.Openid)
	if code != 200 {
		return serializer.ParamErr("登录失败,请再次登录", nil)
	}
	wxinfo := model.Wxuser{
		Openid:     resp.OpenID,
		SessionKey: resp.SessionKey,
		ErrCode:    resp.ErrCode,
		ErrMSG:     resp.ErrMSG,
		AvatarUrl:  service.AvatarUrl,
		Nick:       service.Nick,
		Sex:        service.Sex,
	}

	//检查是否请求成功了

	//创建用户
	//在创建前先判断用户是否已经存在了。
	code = service.ValidUser(resp.OpenID)
	fmt.Println("code", code)

	if code == 200 {
		if err := model.DB.Create(&wxinfo).Error; err != nil {
			return serializer.ParamErr("创建用户失败", err)

		}
	} else {
		//已经存在的用户就把他查询出来
		// var usr model.Wxuser
		err := model.DB.Where("openid = ?", resp.OpenID).First(&wxinfo).Error
		if err != nil {
			return serializer.ParamErr("查询失败", nil)
		}
	}
 
	return serializer.BuildOpenResponse(wxinfo, token)

}
