package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/service"
)

// 获取小程序openid
func GetOpenId(c *gin.Context){
	var service service.GetOpenIdService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Getid()
		c.JSON(200,res)
	}else{
		c.JSON(500,ErrorResponse(err))
	}
}