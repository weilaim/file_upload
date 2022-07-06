package cache

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

func Like(name string, active uint64, likeUserid uint64) (int64, int64, error) {

	//记录点赞
	var like int64
	var isLike int64
	res, err := RedisClient.HMSet(LikeKey, GetFieldActive(name, active)).Result()
	//另外做记录方便查询的时候返回前端
	RedisClient.HMSet(LikeActive, GetActive(likeUserid, active)).Result()
	if active == 1 { //点赞
		if res == "OK" {
			like = 1
			isLike = 1
		}

	} else { //取消点赞
		if res == "OK" {
			like = 0
			isLike = -1
		}

	}

	//count

	count, _ := RedisClient.HIncrBy(LikeKeyCount, FormatActive(likeUserid), isLike).Result()

	return like, count, err
}

func LikeGet(name string, LikeUserId uint64) (int, int, error) {
	//点赞状态
	reActive, _ := RedisClient.HGet(LikeKey, name).Result()

	kecount, err := RedisClient.HGet(LikeKeyCount, strconv.Itoa(int(LikeUserId))).Result()
	// if err != nil {
	// 	return serializer.Err(500, "查询出错", err)
	// }
	isActive, _ := strconv.Atoi(reActive)
	isCount, _ := strconv.Atoi(kecount)

	return isActive, isCount, err
}

/**
 * 点赞。状态为1
 * @param likedUserId
 * @param likedPostId
 */
func SaveLiked2Redis(likedUserId uint64, likedPostId uint64) *redis.StatusCmd {
	name := LoveField(likedUserId, likedPostId)
	StatusLike := RedisClient.HMSet(LikeKey, GetFieldActive(name, 1))
	return StatusLike
}

/**
 * 取消点赞。将状态改变为0
 * @param likedUserId
 * @param likedPostId
 */
func UnlikeFromRedis(likedUserId uint64, likedPostId uint64) *redis.StatusCmd {
	name := LoveField(likedUserId, likedPostId)
	StatusLike := RedisClient.HMSet(LikeKey, GetFieldActive(name, 0))
	return StatusLike
}

/**
 * 从Redis中删除一条点赞数据
 * @param likedUserId
 * @param likedPostId
 */
func DeleteLikedFromRedis(likedUserId uint64, likedPostId uint64) *redis.IntCmd {
	name := LoveField(likedUserId, likedPostId)
	like := RedisClient.HDel(LikeKey, name)
	return like
}

/**
 * 该用户的点赞数加1
 * @param likedUserId count, _ := RedisClient.HIncrBy(LikeKeyCount, FormatActive(likeUserid), isLike).Result()
 */
func IncrementLikedCount(likedUserId uint64) *redis.IntCmd {
	like := RedisClient.HIncrBy(LikeKeyCount, FormatActive(likedUserId), 1)
	return like
}

/**
 * 该用户的点赞数减1
 * @param likedUserId
 */
func DecrementLikedCount(likedUserId uint64) *redis.IntCmd {
	like := RedisClient.HIncrBy(LikeKeyCount, FormatActive(likedUserId), -1)
	return like
}

/**
 * 获取Redis中存储的所有点赞数据
 * @return
 */
type Liked struct {
	LikedUserId string
	LikedPostId string
	LikeStatu   string
}

func GetLikedDataFromRedis() []Liked {
	key, _ := RedisClient.HScan(LikeKey, 0, "*", 0).Val()

	var LikeList []Liked
	var TempId []Liked
	var TempStu []Liked
	var userid string
	var postid string
	var likeStu string
	for _, v := range key {
		if len(v) >= 3 {
			like := strings.Split(v, "::")
			userid = like[0]
			postid = like[1]
			item := Liked{
				LikedUserId: userid,
				LikedPostId: postid,
			}
			TempId = append(TempId, item)
		} else {
			likeStu = v
			item := Liked{
				LikeStatu: likeStu,
			}
			TempStu = append(TempStu, item)
		}

		//存到 list 后从 Redis 中删除
		RedisClient.HDel(LikeKey, v)

	}

	for z, l := range TempId {

		item := Liked{
			LikedUserId: l.LikedUserId,
			LikedPostId: l.LikedPostId,
			LikeStatu:   TempStu[z].LikeStatu,
		}
		fmt.Println(item)
		LikeList = append(LikeList, item)

	}

	return LikeList
}

type AllCount struct {
	LikedPostId string
	LikeSum     string
}

/**
 * 获取Redis中存储的所有点赞数量
 * @return
 */
func GetLikedCountFromRedis() []AllCount {
	key, _ := RedisClient.HScan(LikeKeyCount, 0, "*", 0).Val()

	var LikeListCount []AllCount
	var TempPost []AllCount
	var TempCount []AllCount
	for h, j := range key {

		if h%2 != 0 {
			item := AllCount{
				LikedPostId: j,
			}
			TempPost = append(TempPost, item)
		} else {
			item := AllCount{
				LikeSum: j,
			}

			TempCount = append(TempCount, item)
		}
		//存到 list 后从 Redis 中删除
		RedisClient.HDel(LikeKeyCount, j)

	}

	for n, m := range TempPost {
		count := AllCount{
			LikedPostId: m.LikedPostId,
			LikeSum:     TempCount[n].LikeSum,
		}
		LikeListCount = append(LikeListCount, count)
	}

	return LikeListCount
}
