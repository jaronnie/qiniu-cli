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
	ret := storage.PutRet{}
	formUploader := storage.NewFormUploader(&cfg)
	upToken := "zM_pCg7E1kBvwD0DlGQUadtxVGkumuKNQIDVI4Nl:g-UdVIMkyn6z4uSZO-k51oV69vo=:eyJzY29wZSI6Im5qLWpheSIsImRlYWRsaW5lIjoxNjE2NjkxMjEzfQ=="
	upload := "test.md"
	path := "test.md"
	err := formUploader.PutFile(context.Background(), &ret, upToken, upload, path, &putExtra)
	if err != nil {
		fmt.Println(err)
	}
}
