package global

import (
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"gorm.io/gorm"
)

var (
	MAC *qbox.Mac
	QN_DB *gorm.DB
)
