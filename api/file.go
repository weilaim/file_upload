package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/service"
)


func CreateFile(c *gin.Context){
	service := service.CreateFileService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Files()
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

//ListFiles
func ListFiles(c *gin.Context){
	service := service.ListFilesService{}
	if err := c.ShouldBind(&service);err == nil {
		res := service.List()
		c.JSON(200,res)
	}else{
		c.JSON(500,ErrorResponse(err))
	}
}

//获取单个文件流
func GetFile(c *gin.Context){
	service := service.GetFileService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.SingoFile()
		c.JSON(200,res)
	}else{
	c.JSON(500,ErrorResponse(err))
	}
}
//更新/修改文件流
func UpdateFile(c *gin.Context){
	service := service.UpadteFileService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200,res)
	}else{
		c.JSON(500,ErrorResponse(err))
	}
	
}

func DeleteFile(c *gin.Context){
	service := service.DeleteFileService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200,res)
}