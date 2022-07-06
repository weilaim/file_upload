package model

import "gorm.io/gorm"



type LikeInfoCount struct {
	gorm.Model
	LikePostId string
	LikeCount string
}


// 要是查具体某个(帖子，人)获得的点赞数量，就需要把Redis里的和数据库里的数量都取出来再相加