package service

import (

	"github.com/weilaim/blog-api/cache"
	"github.com/weilaim/blog-api/serializer"
)

//点赞服务   点赞状态为1  取消点赞为2
type CreateLikeService struct {
	LikePostId uint64 `form:"likepostid" json:"likepostid"`                    //被点赞人
	LikeUserId uint64 `form:"likeuserid" json:"likeuserid" binding:"required"` //点赞的人
	LikeActive uint64 `form:"likeactive"  json:"likeactive"`
	LikeGet    string `form:"likeget" json:"likeget"` //查询
}

func (service *CreateLikeService) Like() serializer.Response {

	name := cache.LoveField(service.LikeUserId, service.LikePostId)
	if service.LikeGet != "" {
		isActive, isCount, err := cache.LikeGet(name, service.LikePostId)
		if err != nil {
			return serializer.Response{
				Code: 500,
				Msg:    "操作失败",
			}
		}
		return serializer.Response{
			Code: 200,
			Data: map[string]int{
				"islike":    isActive,
				"likecount": isCount,
			},
		}

	}

	if service.LikeActive == 1 {

		ok, err := cache.SaveLiked2Redis(service.LikeUserId, service.LikePostId).Result()
		if err != nil {
			return serializer.Response{
				Code: 500,
				Msg:    "点赞失败",
				Error:  err.Error(),
			}
		}

		return serializer.Response{
			Code: 200,
			Msg:    "点赞成功" + ok,
			Data: map[string]int64{
				"islike": 1,
				// "likecount": count,
			},
		}
	} else {

		// ok, err := cache.UnlikeFromRedis(service.LikeUserId, service.LikePostId).Result()
		// ok, err := cache.DeleteLikedFromRedis(service.LikeUserId, service.LikePostId).Result() //仅用于测试
		// cache.GetLikedDataFromRedis() //test
		// cache.GetLikedCountFromRedis() //test

 
		 
	 
		// fmt.Println(res)
		// fmt.Println("key",key)
		// fmt.Println("corso",corso)

		// if err != nil {
		// 	return serializer.Response{
		// 		Code: 500,
		// 		Msg:    "点赞失败",
		// 		Error:  err.Error(),
		// 	}
		// }

		return serializer.Response{
			Code: 200,
			// Msg:    "取消点赞成功" + ok,
			Data: map[string]int64{
				"islike": 0,
				// "likecount": count,
			},
		}
	}

}
