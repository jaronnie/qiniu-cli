package fs

import (
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/viper"
)

//bucket

type BucketManger struct {
	Account
	bucket string
	cfg    storage.Config
}

func (bm *BucketManger) GetBucket() string {
	return viper.GetString("qn.bucket")
}
