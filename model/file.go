package model

import "gorm.io/gorm"

type Files struct {
	gorm.Model 
	UserId  uint
	AccName string
	Title   string   
	FileUrl string  
	Content string  `gorm:"type:longText"`	
}