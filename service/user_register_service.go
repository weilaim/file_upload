package service

import (
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=2,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=3,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=3,max=40"`
}


// Valid 验证表单
func (service *UserRegisterService) Valid() *serializer.Response {
	 
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40002,
			Msg:    "两次输入的密码不相同",
		}
	}

	var count int64
	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:    "昵称被占用",
		}
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:    "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Code:   model.Active,
		Avatar: "https://img2.woyaogexing.com/2022/06/20/50a536d88eabca38!400x400.jpg",
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return user, err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return user, &serializer.Response{
			Code: 40002,
			Msg:    "密码加密失败",
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return user, &serializer.Response{
			Code: 40002,
			Msg:    "注册失败",
		}
	}

	return user, nil
}
