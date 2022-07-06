package service

import (
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)


type CreateFileAccService struct{
	Accname  string `form:"accname" json:"accname" binding:"required,min=2"`
	Password   string  `form:"password" json:"password" binding:"required,min=6,max=20" `
	Fieldid	string	`form:"fieldid" json:"fieldid" binding:"required"`
}

// ValidatePassword
func (service *CreateFileAccService) ValidateName() bool {
	var CheckName int64
	model.DB.Model(model.FilesAcc{}).Where("accname = ?", service.Accname).Count(&CheckName)
	if CheckName > 0 {
		return false
	} else {
		return true
	}

}

func (service *CreateFileAccService)FileAcc() serializer.Response{
	Acc := model.FilesAcc{
		Accname: service.Accname,
		Fieldid: service.Fieldid,
	}

	//表单验证
	checkName := service.ValidateName()
	if !checkName {
		return serializer.Response{
			Code: 403,
			Msg: "名字已经存在,请更换",
		}
	}

	//密码加密
	if err := Acc.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Code: 402,
			Msg: "密码加密失败",
		}
	}

	//创建fileAcc
	if err := model.DB.Create(&Acc).Error; err != nil {
		return serializer.Response{
			Code: 402,
			Msg: "添加失败了呢亲！",
		}
	}
	
	return serializer.Response{
		Code: 200,
		Msg: "创建成功啦",
	}
}