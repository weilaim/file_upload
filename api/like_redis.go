package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/service"
)

/**
 * 保存点赞记录
 * @param userLike
 * @return
 */
func Save(c gin.Context) {
	service := service.LikeRedisService{}
	if err :=c.ShouldBind(&service); err == nil {
		res := service.Save()
		c.JSON(200,res)
	}else{
		c.JSON(500,ErrorResponse(err))
	}
}

/**
 * 批量保存或修改
 * @param list
 */
func SaveAll(list []string) {

}

/**
 * 根据被点赞人的id查询点赞列表（即查询都谁给这个人点赞过）
 * @param likedUserId 被点赞人的id
 * @param pageable  //分页
 * @return
 */
func GetLikedListByLikedUserId(likedUserId string, Pageable string) {

}

/**
 * 根据点赞人的id查询点赞列表（即查询这个人都给谁点赞过）
 * @param likedPostId
 * @param pageable //分页
 * @return
 */
func GetLikedListByLikedPostId(likedPostId string, Pageable string) {

}

/**
 * 通过被点赞人和点赞人id查询是否存在点赞记录
 * @param likedUserId
 * @param likedPostId
 * @return
 */
func GetByLikedUserIdAndLikedPostId(likedUserId string, likedPostId string) {

}

/**
 * 将Redis里的点赞数据存入数据库中
 */
func TransLikedFromRedis2DB() {

}

/**
 * 将Redis中的点赞数量数据存入数据库
 */
func TjansLikedCountFromRedis2DB() {

}
