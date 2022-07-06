package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/service"
)

func CreateFileAcc(c *gin.Context) {
	service := service.CreateFileAccService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.FileAcc()
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

//ListFileAcc
func ListFileAccs(c *gin.Context) {
	service := service.ListFilesAccService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

//获取单个Acc文件流
func GetFileAcc(c *gin.Context) {
	service := service.GetFileAccService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.FileAcc(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

//更新/修改文件流
func UpdateFileAcc(c *gin.Context) {
	service := service.UpadteFileService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}

}

func DeleteFileAcc(c *gin.Context) {
	service := service.DeleteFileAccService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}

// 授权得到fileinfo相关信息。
func FileAccAuth(c *gin.Context) {
	service := service.FileAccAuthService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.FileAuth()
		c.JSON(200, res)
	}else{
		c.JSON(500,ErrorResponse(err))
	}

}
