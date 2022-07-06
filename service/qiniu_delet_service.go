package service

import (
	"os"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/weilaim/blog-api/serializer"
)

type QiniuDeletService struct {
	Key string `form:"key" json:"key"`
}

func (service *QiniuDeletService) Del() serializer.Response {
	mac := qbox.NewMac(os.Getenv("QI_ACCESSKEY"), os.Getenv("QI_SECRTKEY"))
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: true, //是否使用cdn加速域名
		UseHTTPS:      true, //是否使用https域名
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	bucket := os.Getenv("QI_BUCKET")

	err := bucketManager.Delete(bucket, service.Key)

	if err != nil {
		return serializer.Response{
			Code: 5000,
			Msg: "删除失败请重试！",
			Error: err.Error(),
		}
		 
	}

	 
	return serializer.Response{
		Code: 200,
		Msg: "删除成功",
	}
}
