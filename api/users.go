package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/service"
)

func ShowUser(c *gin.Context) {
	service := service.ListUserService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}

}

func UserInfo(c *gin.Context) {
	service := service.UserInfoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(404, ErrorResponse(err))
	}

}
