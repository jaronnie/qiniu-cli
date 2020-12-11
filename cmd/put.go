/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (

	overwrite bool

)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put <local file> [flags]",
	Short: "put file to qiniu server",
	Long: `use put command, you can put your local file to qiniu server, and get a url.`,
	Run: put,
}

func put(cmd *cobra.Command, params []string){

	if len(params) <= 0 {

		fmt.Println("please set filepath...")

		os.Exit(0)

	}

	path := params[0]

	index := strings.LastIndexAny(path, "/")

	var upload string

	if index != -1 {

		upload = path[index+1:]

	} else {

		upload = path

	}

	accessKey := viper.GetString("ak")

	secretKey := viper.GetString("sk")

	bucket := viper.GetString("bucket")

	mac := qbox.NewMac(accessKey, secretKey)

	cfg := storage.Config{

		ApiHost:"http://api.qiniu.com",

	}

	var putPolicy storage.PutPolicy

	if overwrite {

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

	ret := storage.PutRet{}

	putExtra := storage.PutExtra{

		Params: map[string]string{

			"x:name": "picture or some other data",
		},
	}


	err = formUploader.PutFile(context.Background(), &ret, upToken, upload, path, &putExtra)

	if err != nil {

		fmt.Println("上传失败")

		return
	}

	fmt.Println("upload successfully")

	fmt.Println("外链为:")

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
