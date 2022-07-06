package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/service"
)

//创建小情书接口
func CreateLove(c *gin.Context) {
	service := service.CreateLoveService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200,res)
	}else{
		c.JSON(500,ErrorResponse(err))
	}

}

//获取小情书列表
func ListLove(c *gin.Context){
	service := service.ListLoveService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200,res)
	}else{
		c.JSON(500,ErrorResponse(err))
	}
}

//获取单个小情书图集数据
func GetLove(c *gin.Context){
	service := service.GetLoveService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Love()
		c.JSON(200,res)
	}else{
		c.JSON(500,ErrorResponse(err))
	}
}


