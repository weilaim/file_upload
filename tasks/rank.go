package tasks

import "github.com/weilaim/blog-api/cache"

// RestartDailyRank 重启一天的排名
func RestartDailyRank() error {
	return cache.RedisClient.Del("rank:daily").Err()
}