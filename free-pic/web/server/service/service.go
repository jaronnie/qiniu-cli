package service

import (
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/viper"
	"qn/server/global"
)

func GenerateToken() string {
	bucket := viper.GetString("qiniu.bucket")
	putPolicy := storage.PutPolicy{Scope: bucket}
	upToken := putPolicy.UploadToken(global.MAC)
	return upToken
}