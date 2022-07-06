package model

import "gorm.io/gorm"

type LikeInfo struct {
	gorm.Model
	LikedUserId string
	LikedPostId string
	LikeStatus  string
	 
}

