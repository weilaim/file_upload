package service

import (
	"fmt"
	// "io"
	"mime"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"github.com/weilaim/blog-api/serializer"
)

// UploadTokenService 获得上传oss token服务

type UploadTokenService struct {
	Filename string `form:"filename" json:"filename"`
}

//Post 创建toke
func (service *UploadTokenService) Post() serializer.Response {
	if service.Filename == "" {
		return serializer.Response{
			Code: 50000,
			Msg: "请选择文件",
		}
	}

	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return serializer.Response{
			Code: 50002,
			Msg:    "oss配置错误",
			Error:  err.Error(),
		}
	}

	// 获取存储空间。
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return serializer.Response{
			Code: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 获取扩展名
	// var filename = "/temp"

	ext := filepath.Ext(service.Filename)
	fmt.Println(service.Filename)
	fmt.Println(ext + "文件扩展")
 

	// 带可选参数的签名直传。
	options := []oss.Option{
		oss.ContentType(mime.TypeByExtension(ext)),
		// oss.ContentType("image/png"),
	}

	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ext
	// err = bucket.PutObjectFromFile(key, localFile)
	// if err != nil {
	// 	HandleError(err)
	// }
	// 签名直传。
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	// signedPutURL, err := bucket.SignURL(key, oss.HTTPPost, 600, options...)
	if err != nil {
		return serializer.Response{
			Code: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}
	// 查看图片
	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		return serializer.Response{
			Code: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			"get": signedGetURL,
		},
	}
}
