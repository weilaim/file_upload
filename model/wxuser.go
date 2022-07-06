package model

import "gorm.io/gorm"

type Wxuser struct {
	gorm.Model
	Id int64
	Code string  
	Nick string 
	AvatarUrl string  
	Sex int32  
	Openid string
	SessionKey string
	UnionId string	
	ErrCode int
	ErrMSG string
}