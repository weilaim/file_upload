package service

import (

	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/middleware"
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

//UserLoginService 管理用户登录服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=2,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=40"`
}

//用户登录 Login
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.Response{
			Code:402 ,
			Msg: "账号或密码错误",
			Error: err.Error(),
		}
	}

	if !user.CheckPassword(service.Password) {
		return serializer.Response{
			Code:402 ,
			Msg: "账号或密码错误",
		}
	}
	//设置token
	token, code := middleware.SetToken(service.UserName)
	if code != 200 {
		return serializer.ParamErr("登录失败,请再次登录", nil)
	}
	return serializer.BuildWxuserResponse(user, token, "登陆成功")
}
