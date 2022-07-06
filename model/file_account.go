package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type FilesAcc struct {
	gorm.Model 
	Accname  string
	Password   string   
	Fieldid	string
}


//SetPassword 设置密码
func (file *FilesAcc) SetPassword(password string) error{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),PassWordConst)
	if err != nil{
		return err
	}

	file.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func(file *FilesAcc) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(file.Password),[]byte(password))
	return err == nil
}
