package model

//执行自动迁移

func migration(){
	//自动迁移模式
	_ = DB.AutoMigrate(&User{})
	//自动迁移Video表
	_ = DB.AutoMigrate(&Video{})

	//自动迁移Tarp表
	_ = DB.AutoMigrate(&Wxuser{})

	//图片图集
	_ = DB.AutoMigrate(&Loveimg{})

	//表白表
	_ = DB.AutoMigrate(&Loves{})

	//文件管理表
	_ = DB.AutoMigrate(&Files{})

	//FileAcc
	_ = DB.AutoMigrate((&FilesAcc{}))
}