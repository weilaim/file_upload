package service

import "github.com/weilaim/blog-api/serializer"

type LikeRedisService struct{

}

/**操作方法
 * 保存点赞记录
 * @param userLike
 * @return
 */
func(s *LikeRedisService) Save() serializer.Response  {
	//从redis中获取点赞记录 
	// return likeRepository.save(userLike);
	return serializer.Response{

	}
}
/**操作方法
 * 批量保存或修改
 * @param list
 */
// func(s *LikeRedisService)  saveAll( list []string) {
// 	return likeRepository.saveAll(list);
// }

 /** -------------应该是需要返回前端----------
 * 根据被点赞人的id查询点赞列表（即查询都谁给这个人点赞过）   
 * @param likedUserId 被点赞人的id
 * @param pageable  //分页
 * @return
 */
// func(s *LikeRedisService) getLikedListByLikedUserId(String likedUserId, Pageable pageable) {
// 	return likeRepository.findByLikedUserIdAndStatus(likedUserId, LikedStatusEnum.LIKE.getCode(), pageable);
// }

 /**-------------应该是需要返回前端----------
 * 根据点赞人的id查询点赞列表（即查询这个人都给谁点赞过）
 * @param likedPostId
 * @param pageable //分页
 * @return
 */
//  func(s *LikeRedisService) getLikedListByLikedPostId(String likedPostId, Pageable pageable) {
// 	return likeRepository.findByLikedPostIdAndStatus(likedPostId, LikedStatusEnum.LIKE.getCode(), pageable);
// }

 /** 操作方法
 * 通过被点赞人和点赞人id查询是否存在点赞记录
 * @param likedUserId
 * @param likedPostId
 * @return
 */
// func(s *LikeRedisService)  getByLikedUserIdAndLikedPostId(String likedUserId, String likedPostId) {
// 	return likeRepository.findByLikedUserIdAndLikedPostId(likedUserId, likedPostId);
// }

 /**-------------应该是需要定时执行----------
 * 将Redis里的点赞数据存入数据库中
 */
//  func(s *LikeRedisService) transLikedFromRedis2DB() {
// 	List<UserLike> list = redisService.getLikedDataFromRedis();
// 	for (UserLike like : list) {
// 		UserLike ul = getByLikedUserIdAndLikedPostId(like.getLikedUserId(), like.getLikedPostId());
// 		if (ul == null){
// 			//没有记录，直接存入
// 			save(like);
// 		}else{
// 			//有记录，需要更新
// 			ul.setStatus(like.getStatus());
// 			save(ul);
// 		}
// 	}
// }

 /**-------------应该是需要定时执行----------
 * 将Redis中的点赞数量数据存入数据库
 */
//  func(s *LikeRedisService)  transLikedCountFromRedis2DB() {
// 	List<LikedCountDTO> list = redisService.getLikedCountFromRedis();
// 	for (LikedCountDTO dto : list) {
// 		UserInfo user = userService.findById(dto.getId());
// 		//点赞数量属于无关紧要的操作，出错无需抛异常
// 		if (user != null){
// 			Integer likeNum = user.getLikeNum() + dto.getCount();
// 			user.setLikeNum(likeNum);
// 			//更新点赞数量
// 			userService.updateInfo(user);
// 		}
// 	}
// }
