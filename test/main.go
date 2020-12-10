package main

import (
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

func main() {


	accessKey := "0BXxOytiWWOySUiEbg-t8_08c0tvPpwTwDA6Ivfn"

	secretKey := "ewA76OHGBEz43vQg-gCqOHd_5paEurmSEzFdx-dz"

	bucket := "nj-jay"

	mac := qbox.NewMac(accessKey, secretKey)

	cfg := storage.Config{

		ApiHost: "api.qiniu.com",
	}

	bm := storage.NewBucketManager(mac, &cfg)

	domains, err := bm.ListBucketDomains(bucket)

	if err != nil {

		fmt.Println(err)

	} else {

		fmt.Println(domains)
	}


}
