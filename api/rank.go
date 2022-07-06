package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/service"
)

// DailyRank 每日排行榜
func DailyRank(c *gin.Context){ 
	service := service.DalyRankService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get()
		c.JSON(200,res)
	}

}