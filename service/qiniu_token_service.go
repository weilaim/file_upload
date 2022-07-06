package service

import (
	"context"

	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/weilaim/blog-api/serializer"
)

type QiniuTokenService struct {
}

func (service *QiniuTokenService) Post(file multipart.File, fileSize int64, filename string) serializer.Response {
	//获取文件扩展名
 
	ext := filepath.Ext(filename)
	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ext

	putPolicy := storage.PutPolicy{
		Scope: os.Getenv("QI_BUCKET"),
	}
	mac := qbox.NewMac(os.Getenv("QI_ACCESSKEY"), os.Getenv("QI_SECRTKEY"))
	upToken := putPolicy.UploadToken(mac)

	cfq := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: true, //是否使用cdn加速域名
		UseHTTPS:      true, //是否使用https域名
	}
	putExtra := storage.PutExtra{}                //key 可能在这里              // PutExtra 为表单上传的额外可选项
	formUploader := storage.NewFormUploader(&cfq) //构建一个表单上传的对象
	ret := storage.PutRet{}                       //PutRet 为七牛标准的上传回复内容。

	// // err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	// // err := formUploader.PutFile(context.Background(), &ret, upToken, key, file, &putExtra)
	err := formUploader.Put(context.Background(), &ret, upToken, key, file, fileSize, &putExtra)
	if err != nil {
		return serializer.Response{
			Code: 500,
			Errno: 2,
			Message: "上传错误",
			Msg:    "上传错误",
			Error:  err.Error(),
		}
	}

	url := os.Getenv("QI_SERVER") + ret.Key

	return serializer.Response{
		Message: "hahaha",
		Errno: 0,
		Code: 200,
		Data: map[string]string{
			// "url": "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg.jj20.com%2Fup%2Fallimg%2F1113%2F052420110515%2F200524110515-2-1200.jpg&refer=http%3A%2F%2Fimg.jj20.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1652435739&t=f8a8dffded60de70cf258282e9173fd3",
			"url": url,
			"key": ret.Key,
		},
	}
}
