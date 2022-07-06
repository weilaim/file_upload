package cache

import (
	"fmt"
	"strconv"
)

const (
	// DailyRankKdy 每日排行
	DailyRankKdy = "rank:daily"
	//点赞key
	LikeKey      = "MAP_USER_LIKED"
	LikeActive = "MAP_USER_ACTIVE"
	LikeKeyCount = "MAP_KEY_USER_LIKED_COUNT"
	//收藏key
	CollectionKey = "USER_COLLEC"
	CollectionCount = "USER_COLLEC_COUNT"
	MyCollec = "MY_COLLECID"

)

// VideoViewKey 点击视频的key
//view:video：1 -》 100
//view:video: 2 -> 150
func LoveViewKey(id uint) string {
	return fmt.Sprintf("view:love:%s", strconv.Itoa(int(id)))
}

//likedUserId::likedPostId
func LoveField(LikeUserId uint64, LikePostId uint64) string {
	Lid := FormatActive(LikeUserId)
	Lpostid := FormatActive(LikePostId)
	return fmt.Sprintf("%s::%s",Lid , Lpostid)

}


/**
 * @description: 
 * @param {uint64} LikeUserId
 * @return {string} "LikeuserID::*"
 */
func LikeMatch(LikeUserId uint64) string {
	likeid := FormatActive(LikeUserId)
	return fmt.Sprintf("%s::",likeid)
}

//CollecUserId::CollectPostId
func CollecField(CollecId uint64, CollecPostId uint64) string {
	Cid := FormatActive(CollecId)
	Cpostid := FormatActive(CollecPostId)
	return fmt.Sprintf("%s::%s", Cid, Cpostid)
	
}

//collection 收藏key

/**
 * @description: 
 * @param {uint64} id
 * @return {*}
 */
func FormatActive(id uint64) string {
	return strconv.FormatUint(id, 10)
}


/**
 * @description: 
 * @param {string} name
 * @param {uint64} active
 * @return {map[string]interface}
 */
func GetFieldActive(name string, active uint64) map[string]interface{} {
	field := map[string]interface{}{
		name: active,
	}
	return field
}

func GetActive(name uint64,active uint64) map[string]interface{}{
	key := strconv.FormatUint(name,10)
	field := map[string]interface{}{
		key:active,
	}
	return field
}