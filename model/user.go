package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//User 用户模型

type User struct{
	gorm.Model
	UserName string
	PasswordDigest string
	Nickname string
	Code string
	Avatar string `gorm:"size:1000"`
}

const (
	//PassWordConst 密码难度
	PassWordConst = 12
	//Active 激活用户
	Active string = "active"
	//Inactive 未激活用户
	Inactive string = "inactive"
	//Suspend 被封禁用户
	Suspend string = "suspend"
)

//获取用户id
func GetUser(ID interface{}) (User, error){
	var user User
	result := DB.First(&user,ID)
	return user,result.Error
}

//SetPassword 设置密码
func (user *User) SetPassword(password string) error{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),PassWordConst)
	if err != nil{
		return err
	}

	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func(user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest),[]byte(password))
	return err == nil
}
