package service

import (
	"fmt"
	"strconv"

	"github.com/weilaim/blog-api/middleware"
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

type FileAccAuthService struct {
	Password string `form:"password" json:"password" binding:"required,min=6"`
	FileId   uint   `form:"file_id" json:"file_id"`
}

//请求穿过来的Accid 去查询单个fileinfo的信息返回去
func (service *FileAccAuthService) FileAuth() serializer.Response {
	var accModel model.FilesAcc
	fmt.Println(service.FileId)
	if err := model.DB.Where("id = ?", service.FileId).First(&accModel).Error; err != nil {
		return serializer.Response{
			Code:  402,
			Msg:   "输入有误",
			Error: err.Error(),
		}
	}
	if !accModel.CheckPassword(service.Password) {
		return serializer.Response{
			Code: 402,
			Msg:  "密码错误",
		}
	}
	fmt.Println(service.Password)
	//设置token
	token, code := middleware.SetToken(service.Password)
	if code != 200 {
		return serializer.ParamErr("登录失败,请再次登录", nil)
	}
	return serializer.Response{
		Code: 200,
		Data: map[string]string{
			"token": token,
			"id":    strconv.Itoa(int(accModel.ID)),
		},
	}
}
