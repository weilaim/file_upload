package service

import (
	"github.com/weilaim/blog-api/cache"
	"github.com/weilaim/blog-api/serializer"
)

//collection 收藏服务 收藏状态为1  取消状态为2
type CreateCollection struct {
	CId      uint64 `form:"cid" json:"cid"`         //被收藏id
	CpostId  uint64 `form:"cpostid" json:"cpostid"` //收藏人
	Cactive  uint64 `form:"cactive" json:"cactive"`
	ColleGet string `form:"colleget" json:"colleget"` //查询
}

var ColleActive uint64 = 1 // 点赞中的状态
/**
 * @description: 
 * @param {*}
 * @return {*}
 */
func (service *CreateCollection) Collec() serializer.Response {
	name := cache.CollecField(service.CId, service.CpostId)
	if service.Cactive == 1 {
		//记录收藏人和收藏love.id 信息
		//记录是谁收藏了那些love.id key 就是userid   fields love.id
		cache.RedisClient.HMSet(cache.MyCollec, cache.GetActive(service.CpostId, service.CId)).Result()
		//收藏的状态
		cCode, count, err := cache.SetCollection(name, service.Cactive, service.CpostId)
		if err != nil {
			return serializer.Response{
				Code: 500,
				Msg:    "操作失败",
				Error:  err.Error(),
			}
		}

		return serializer.Response{
			Code: 200,
			Msg:    "收藏成功",
			Data: map[string]int64{
				"iscollec": cCode,
				"collesum": count,
			},
		}

	} else {
		//取消收藏 就应该删除记录cache.GetActive(service.CpostId, service.CId)
		//记录是谁收藏了那些love.id key 就是userid   fields love.id
		cache.RedisClient.HDel(cache.MyCollec,).Result()
		//取消收藏
		ColleActive = 0
		uncolle, count, err := cache.SetCollection(name, ColleActive, service.CpostId)
		if err != nil {
			return serializer.Response{
				Code: 500,
				Msg:    "操作失败",
				Error:  err.Error(),
			}
		}

		return serializer.Response{
			Code: 200,
			Msg:    "取消成功",
			Data: map[string]int64{
				"iscollec": uncolle,
				"collesum": count,
			},
		}

	}

}
