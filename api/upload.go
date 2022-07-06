package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/service"
)

// UploadToken 上传授权
func UploadToken(c *gin.Context) {
	service := service.UploadTokenService{}
	if err := c.ShouldBind(&service); err == nil {
		// file, fileHeader, _ := c.Request.FormFile("filename")
		// if file != nil {
		// 	fmt.Println(fileHeader.Header)
		// 	res := service.Post(file, fileHeader.Filename)
		// 	c.JSON(200, res)
		// } else {
		// 	c.JSON(500, ErrorResponse(err))
		// 	return
		// }
		res := service.Post()

		c.JSON(200, res)

	} else {
		c.JSON(500, ErrorResponse(err))
	}

}
