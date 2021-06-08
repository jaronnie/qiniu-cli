package fs

import (
	"github.com/qiniu/go-sdk/v7/storage"
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

func (bm *BucketManger) NewBucketManger() *storage.BucketManager {
	var acc Account
	return &storage.BucketManager{
		Mac: acc.GetMac(),
		Cfg: &storage.Config{},
	}
}
