package model

import "gorm.io/gorm"

//图集列表
type Loveimg struct {
	gorm.Model
	LovesId  int
	LoveUrl string
}
