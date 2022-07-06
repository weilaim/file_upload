package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/service"
)

//上传文件
func QiniuToken(c *gin.Context){
	service := service.QiniuTokenService{}
	if err := c.ShouldBind(&service); err == nil {
		file,fileHeader,_ := c.Request.FormFile("filename")
		fileSize := fileHeader.Size
		fileName := fileHeader.Filename
		res := service.Post(file,fileSize,fileName)
		c.JSON(200,res)
	}else{
		c.JSON(500,ErrorResponse(err))
	}

}

//删除文件
func QiniuDelet(c *gin.Context){
	service := service.QiniuDeletService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Del()
		c.JSON(200,res)
	}else{
		c.JSON(500,ErrorResponse(err))
	}
}