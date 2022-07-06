package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/service"
)

func CreateLike(c *gin.Context) {
	service := service.CreateLikeService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Like()
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}
