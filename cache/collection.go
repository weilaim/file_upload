package cache

var colleStatu int64
var isColle int64

//记录收藏
func SetCollection(name string, active uint64, colleUser uint64) (int64, int64, error) {
	// CollectionCount
	res, err := RedisClient.HMSet(CollectionKey, GetFieldActive(name, active)).Result()

	if active == 1 {
		//收藏
		if res == "OK" {
			colleStatu = 1
			isColle = 1

		}
	} else {
		//取消收藏
		if res == "OK" {
			colleStatu = 0
			isColle = -1
			//取消了收藏应该删掉记录 那些love.id key 就是userid   fields love.id
			// RedisClient.HMSet(MyCollec, GetActive(colleUser, 1)).Result()
		}
	}

	//count
	colleCount, _ := RedisClient.HIncrBy(CollectionCount, FormatActive(colleUser), isColle).Result()

	return colleStatu, colleCount, err

}
