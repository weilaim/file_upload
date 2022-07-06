package api

import (

	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/service"
)


func Collection(c *gin.Context){
	service := service.CreateCollection{}
	if err := c.ShouldBind(&service);err == nil {
		res := service.Collec()
		c.JSON(200,res)
	}else{
		c.JSON(500,ErrorResponse(err))
	}
}