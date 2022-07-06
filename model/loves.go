package model

import (
	"strconv"

	"github.com/weilaim/blog-api/cache"
	"gorm.io/gorm"
)

// 表白小情书图集
type Loves struct {
	gorm.Model
	Uuid         string
	Uid          int
	Nick         string
	Sex          int
	Content      string
	Lve          string
	Professional string
	Imgsid       int
	Avatar       string
	Loveimg      []Loveimg `gorm:"FOREIGNKEY:LovesId;ASSOCIATION_FOREIGNKEY:ID"`
}

func (Love *Loves) AddView() {
	//增加小情书点击数
	cache.RedisClient.Incr(cache.LoveViewKey(Love.ID))
}

//点击数的获取
func (Love *Loves) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.LoveViewKey(Love.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	
	return count
}

// func formatActive(id uint64) string {
// 	return strconv.FormatUint(id, 10)
// }

// //获取点赞状态
// func GetLike(Love *Loves){
// 	name := cache.LoveField(formatActive(LikeUserId), formatActive(service.LikePostId))
// 	res, err := cache.RedisClient.HGet(cache.LikeKey, name).Result()
// }

// // 获取点赞总数 
func (Love *Loves) LikeView() int {
	kecount, _ := cache.RedisClient.HGet(cache.LikeKeyCount, strconv.Itoa(int(Love.ID))).Result()
	count, _ := strconv.Atoi(kecount)
	return count
}

//获取点赞状态  //先查redis 查不到再查mysql
func (Love *Loves) LikeSatus(uid int) int {
	// name := strconv.Itoa(int(Love.ID))
	name := cache.LoveField(uint64(uid), uint64(Love.ID))
	reActive, _ := cache.RedisClient.HGet(cache.LikeKey, name).Result()
	likestatu, _ := strconv.Atoi(reActive)
	return likestatu
}

//收藏的状态
/**
 * @description:
 * @param {int} uid
 * @return {*}
 */
func (Love *Loves) Collestu(uid int) int {
	// name := strconv.Itoa(int(Love.ID))
	name := cache.CollecField(uint64(uid), uint64(Love.ID))
	lolle, _ := cache.RedisClient.HGet(cache.CollectionKey, name).Result()
	status, _ := strconv.Atoi(lolle)
	return status
}

/**
 * @description:
 * @param {*}
 * @return {*}
 */
func (Love *Loves) ColleCount() int {
	colle, _ := cache.RedisClient.HGet(cache.CollectionCount, strconv.Itoa(int(Love.ID))).Result()
	count, _ := strconv.Atoi(colle)
	return count
}
