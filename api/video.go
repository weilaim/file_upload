package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/service"
)

//CreateVideo 创建视频接口
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// ShowVideo 获取视频详情
func ShowVideo(c *gin.Context) {
	service := service.ShowVideoService{}
	res := service.Show(c.Param("id"))
	c.JSON(200,res)
}

//ListVideo 获取视频列表
func ListVideo(c *gin.Context) {
	service := service.ListVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// UpdateVideo 更新视频
func UpdateVideo(c *gin.Context){
	service := service.UpdateVideoService{}
	if err := c.ShouldBind(&service); err == nil{
		res := service.Update(c.Param("id"))
		c.JSON(200,res)
	}else{
		c.JSON(500,ErrorResponse(err))
	}
}


// DeleteVideo 删除视频
func DeleteVideo(c *gin.Context){
	service := service.DeleteVideoService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200,res)
}