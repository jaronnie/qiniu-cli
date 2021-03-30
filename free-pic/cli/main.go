package main
import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/storage"
)
func main() {
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "picture or some other data",
		},
	}
	cfg := storage.Config{
		ApiHost:"http://api.qiniu.com",
	}
	cfg.Zone = &storage.ZoneHuanan
	ret := storage.PutRet{}
	formUploader := storage.NewFormUploader(&cfg)
	err := formUploader.PutFile(context.Background(), &ret, upToken, upload, path, &putExtra)
	if err != nil {
		fmt.Println(err)
	}
}
