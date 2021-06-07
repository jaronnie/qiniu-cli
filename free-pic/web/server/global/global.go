package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"gorm.io/gorm"
)

var (
	MAC *qbox.Mac
	QN_DB *gorm.DB
	RANK_DB *redis.Client
)
