package service

import (
	"fmt"
	"strings"

	"github.com/weilaim/blog-api/cache"
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

// DalyRankService  每日排行榜服务
type DalyRankService struct {

}

// Get 获取排行
func (service *DalyRankService) Get() serializer.Response {
	var videos []model.Video

	//从redis读取点击前十的视频
	vids, _ := cache.RedisClient.ZRevRange(cache.DailyRankKdy,0,9).Result()

	if len(vids) > 1 {
		order := fmt.Sprintf("FIELD(id,%s)",strings.Join(vids,","))
		err := model.DB.Where("id in (?)",vids).Order(order).Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Code: 50000,
				Msg: "数据库连接错误",
				Error: err.Error(),
			}
		}
	}
	// 返回前十排行列表
	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}

