package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/cobra"
)

var (
	overwrite bool
)

type PutRet struct {
	Hash         string `json:"hash"`
	PersistentID string `json:"persistentId"`
	Key          string `json:"key"`
	Fsize        int    `json:"fsize"`
}

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put <local/remote file> [<local/remote file>] [flags]",
	Short: "put file to qiniu server",
	Long:  `use put command, you can put your local file to qiniu server, and get a url.`,
	Run:   put,
}

func put(cmd *cobra.Command, params []string) {
	if len(params) <= 0 {
		fmt.Println("please set filepath...")
		os.Exit(0)
	}

	if !Acc.ValidAccount() {
		fmt.Println("请设置ak, sk, bucket")
		os.Exit(0)
	}

	mac := Acc.GetMac()
	bucket := Bm.GetBucket()
	cfg := storage.Config{
		ApiHost: "http://api.qiniu.com",
	}

	for _, path := range params {
		if isRemoteFile(path) {
			putRemoteFile(path, mac, bucket, cfg)
		} else {
			putLocalFile(path, mac, bucket, cfg)
		}
	}
}

func isRemoteFile(path string) bool {
	return strings.HasPrefix(path, "http")
}

func getFinalKey(path string) string {
	var finalKey string
	index := strings.LastIndex(path, "/")
	if index != -1 {
		finalKey = path[index+1:]
	}
	return finalKey
}

func putRemoteFile(path string, mac *qbox.Mac, bucket string, cfg storage.Config) {
	finalKey := getFinalKey(path)
	bm := storage.NewBucketManager(mac, &cfg)
	res, err := bm.Fetch(path, bucket, finalKey)
	if err != nil {
		fmt.Println(err)
	}
	domains, _ := bm.ListBucketDomains(bucket)
	fmt.Println("http://" + domains[0].Domain + "/" + res.Key)
}

func putLocalFile(path string, mac *qbox.Mac, bucket string, cfg storage.Config) {
	path = strings.Replace(path, `\`, `/`, -1)
	index := strings.LastIndexAny(path, "/")
	var upload string
	if index != -1 {
		upload = path[index+1:]
	} else {
		upload = path
	}

	var putPolicy storage.PutPolicy

	if overwrite {
		if path == upload {
			err := os.Chdir(".")
			if err != nil {
				fmt.Println("err")
			}
		} else {
			err := os.Chdir(path[0 : index+1])
			path = upload
			if err != nil {
				fmt.Println("err")
			}
		}
		putPolicy = storage.PutPolicy{
			Scope: fmt.Sprintf("%s:%s", bucket, path),
		}
	} else {
		putPolicy = storage.PutPolicy{
			Scope: bucket,
		}

	}
	upToken := putPolicy.UploadToken(mac)
	bm := storage.NewBucketManager(mac, &cfg)
	domains, err := bm.ListBucketDomains(bucket)
	if err != nil {
		fmt.Println("get domain err")
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutPolicy{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "picture or some other data",
		},
	}
	err = formUploader.PutFile(context.Background(), &ret, upToken, upload, path, &putExtra)
	if err != nil {
		fmt.Println(err)
		fmt.Println("上传失败")
		return
	}
	fmt.Println("http://" + domains[0].Domain + "/" + upload)
}

func init() {
	putCmd.Flags().BoolVarP(
		&overwrite,
		"overwrite",
		"w",
		false,
		"when you use -w options, you can replace the same file...",
	)
	rootCmd.AddCommand(putCmd)
}
