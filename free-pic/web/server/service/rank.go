package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"qn/server/global"
)

var ctx = context.Background()


func GetRank(member string) []redis.Z {
	global.RANK_DB.ZIncrBy(ctx, "rank", 1, member)
	ranks, _ := global.RANK_DB.ZRevRangeWithScores(ctx, "rank", 0, -1).Result()
	return ranks
}

