package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"io/ioutil"
)

func main() {
	r := gin.Default()
	r.POST("/qiniu/upload/callback", Callback)
	r.GET("/qiniu/upload/token", RetCliToken)
	r.Run(":9001")
}

func Callback(c *gin.Context) {
	callback, _ := ioutil.ReadAll(c.Request.Body)
	c.IndentedJSON(200, gin.H{
		"data": string(callback),
	})
}


func RetCliToken(c *gin.Context) {
	token := generateToken()
	c.IndentedJSON(200, gin.H{
		"token": token,
	})
}

func generateToken() string {
	ak := "zM_pCg7E1kBvwD0DlGQUadtxVGkumuKNQIDVI4Nl"
	sk := "P2mdTgHaJ7PPX7cggzpBHti-VKppSh-Mqkj-hWeO"
	bucket := "nj-jay"
	mac := qbox.NewMac(ak, sk)
	putPolicy := storage.PutPolicy{Scope: bucket}
	upToken := putPolicy.UploadToken(mac)
	return upToken
}